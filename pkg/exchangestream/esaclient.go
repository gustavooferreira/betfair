package exchangestream

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gustavooferreira/betfair/pkg/globals"
	"github.com/gustavooferreira/betfair/pkg/utils/log"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

const maxBufCapacity = 1024 * 1024

// timeout of 0.5 seconds
const timeoutDuration = 500 * time.Millisecond

// ESAClient is the client that interacts with betfair Exchange Stream API
// It's thread safe!
type ESAClient struct {
	// Treated as immutable
	appKey       string
	sessionToken string

	// Connection to server
	conn *tls.Conn

	// Atomic
	msgID        uint32
	connectionID atomic.Value
	// Wait time (in seconds) before timeout
	connWaitTime uint32
	// When sending or receiving messages through channels, how long to wait before giving up (in seconds)
	chanWaitTime uint32

	reqMsgChan chan WorkUnit

	// Life cycle management
	stopChan       chan bool
	stopInformChan chan bool

	// Change Streams
	MCMChan chan MarketChangeM
	OCMChan chan OrderChangeM
}

func NewESAClient(appKey string, sessionToken string) ESAClient {
	client := ESAClient{appKey: appKey, sessionToken: sessionToken}
	// Set some defaults
	client.connWaitTime = 3
	client.chanWaitTime = 3

	client.connectionID.Store("")
	return client
}

func (esaclient *ESAClient) ChangeSettings(connWaitTime uint32, chanWaitTime uint32) {
	atomic.StoreUint32(&esaclient.connWaitTime, connWaitTime)
	atomic.StoreUint32(&esaclient.chanWaitTime, chanWaitTime)
}

func (esaclient *ESAClient) GetSessionInfo() (string, string, string, uint32) {
	connID := esaclient.connectionID.Load().(string)

	return esaclient.appKey, esaclient.sessionToken, connID, esaclient.msgID
}

// Connect connects to the server.
// When connected, spawn 3 goroutines: controller, reader and writer
func (esaclient *ESAClient) Connect(serverHost string, serverPort uint, insecureSkipVerify bool) (err error) {

	// If connectionID != "" then there is a connection already!
	if esaclient.connectionID.Load().(string) != "" {
		return ConnectionError{Msg: "connection already established"}
	}

	config := tls.Config{InsecureSkipVerify: insecureSkipVerify}
	d := net.Dialer{Timeout: time.Duration(esaclient.connWaitTime) * time.Second}
	esaclient.conn, err = tls.DialWithDialer(&d, "tcp", serverHost+":"+strconv.Itoa(int(serverPort)), &config)
	if err != nil {
		return ConnectionError{Msg: "connecting to betfair failed", Err: err}
	}

	log.Log(globals.Logger, log.INFO, "connection established with server", nil)

	// Init channels
	esaclient.stopChan = make(chan bool)
	esaclient.stopInformChan = make(chan bool)
	esaclient.reqMsgChan = make(chan WorkUnit, 1000)
	esaclient.MCMChan = make(chan MarketChangeM, 1000)
	esaclient.OCMChan = make(chan OrderChangeM, 1000)

	connMsgChan := make(chan ConnectionMessage)

	// Spawn controller in a goroutine
	go esaclient.controller(connMsgChan)

	// Select connMsgChan and Timeout of X seconds
	// If timeout, bring all 3 goroutines down: controller, reader, writer
	select {
	case connMsg := <-connMsgChan:
		esaclient.connectionID.Store(connMsg.ConnectionID)
		return nil

	case <-time.After(time.Duration(esaclient.chanWaitTime) * time.Second):
		// call timed out
		esaclient.disconnectHelper()

		return errors.New("Timeout while waiting for connection message from betfair")
	}
}

// Disconnect disconnects from the server
func (esaclient *ESAClient) Disconnect() error {
	// Check that there is a connection to disconnect
	if esaclient.connectionID.Load().(string) == "" {
		return fmt.Errorf("no connection available to disconnect")
	}

	return esaclient.disconnectHelper()
}

func (esaclient *ESAClient) disconnectHelper() error {
	close(esaclient.stopChan)

	// Wait for the stop inform to arrive or after X seconds kill the connection anyway
	select {
	case <-esaclient.stopInformChan:
		// success tearing down
	case <-time.After(time.Duration(esaclient.chanWaitTime) * time.Second):
		// call timed out
	}

	err := esaclient.conn.Close()
	_ = err

	// Close stream channels as well

	esaclient.connectionID.Store("")
	atomic.StoreUint32(&esaclient.msgID, 0)
	return nil
}

func (esaclient *ESAClient) controller(connMsgChan chan<- ConnectionMessage) {
	fmt.Println("starting controller goroutine")
	defer fmt.Println("exiting controller goroutine")

	connPhaseDone := false

	respMsgChan := make(chan ResponseMessage, 1000)
	reqMsgChan := make(chan RequestMessage, 1000)
	readerStopChan := make(chan bool)
	writerStopInformChan := make(chan bool)

	// LookupTable
	lookupTable := make(map[uint32](chan ResponseMessage))

	// Spawn reader and writer goroutines
	go esaclient.reader(respMsgChan, readerStopChan)
	go esaclient.writer(reqMsgChan, writerStopInformChan)

	// If stopChan is closed, trigger reader and writer to end too
	// When getting the connection message, send the data and close the channel
	for {
		select {
		case _, ok := <-esaclient.stopChan:
			if !ok {
				esaclient.stopChan = nil
				close(readerStopChan)
			}
		case respMsg, ok := <-respMsgChan:
			if !ok {
				close(esaclient.stopInformChan)
				return
			}

			if respMsg.Op == "connection" {
				if !connPhaseDone {
					connMsgChan <- *respMsg.ConnectionMessage
					close(connMsgChan)
					connPhaseDone = true
				} else {
					fmt.Printf("this should have never happened - for another ConnectionMessage: %+v", respMsg)
				}
			} else if respMsg.Op == "mcm" {
				esaclient.MCMChan <- MarketChangeM{ID: respMsg.ID, MarketChangeMessage: *respMsg.MarketChangeMessage}
			} else if respMsg.Op == "ocm" {
				esaclient.OCMChan <- OrderChangeM{ID: respMsg.ID, OrderChangeMessage: *respMsg.OrderChangeMessage}
			} else if respMsg.Op == "status" {
				// Do a lookup and match based on ID!
				if result, ok := lookupTable[respMsg.ID]; !ok {
					// Error, no ID found!
				} else {
					result <- respMsg

					// Delete entry from the lookup table!
					delete(lookupTable, respMsg.ID)
				}
			} else {
				// Error!
			}

			fmt.Printf("Message: %+v\n", respMsg)
		case workUnit, ok := <-esaclient.reqMsgChan:
			if !ok {
				// Error
			}

			// Check if ID is set, if not, get one and set it on the struct
			if workUnit.req.ID == 0 {
				workUnit.req.ID = esaclient.getNewID()
			}

			lookupTable[workUnit.req.ID] = workUnit.respChan

			// Store Id and channel in lookup table
			fmt.Printf("%+v\n", workUnit)
		}
	}
}

// reader is responsible for reading all incoming messages and sending the corresponding objects down the channel
func (esaclient *ESAClient) reader(respMsgChan chan<- ResponseMessage, stopChan <-chan bool) {
	fmt.Println("starting reader goroutine")
	defer fmt.Println("exiting reader goroutine")

	// Create a 8MB buffer for incoming messages
	var buf [8 * 1024 * 1024]byte
	indiceStart := 0
	indiceStop := 0

	for {
		// check stopChan without blocking, if set then exit function
		select {
		case <-stopChan:
			close(respMsgChan)
			return
		default:
		}

		// Call Read with timeout
		esaclient.conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		n, err := esaclient.conn.Read(buf[indiceStop : len(buf)-1])
		log.Log(globals.Logger, log.INFO, fmt.Sprintf("read %d bytes from connection", n), nil)
		if err1, ok := err.(*net.OpError); ok {
			// If timeout, continue
			if err1.Timeout() {
				log.Log(globals.Logger, log.INFO, fmt.Sprintf("timeout %+v", err1), nil)
				continue
			}
			fmt.Printf("ERROR: %T - %+[1]v\n", err)
			// Before continue, prob need to do something with the potential stuff that is in the buffer!
			// TODO: get the extra data read and update indices!
			continue
		} else if err != nil {
			// If EOF, connection was closed!!
			fmt.Printf("ERROR: %T - %+[1]v\n", err)
			// Before continue, prob need to do something with the potential stuff that is in the buffer!
			// TODO: get the extra data read and update indices!
			continue
		}

		indiceStop += n
		progress := true

		// Iterate through the various potential messages (delim \r\n) we might have
		for progress {
			progress = false

			for i := indiceStart; i < indiceStop; i++ {
				if buf[i] == '\n' {
					if i-indiceStart == 0 || i-indiceStart == 1 {
						// discard
						indiceStart = i + 1
						progress = true
						break
					} else if buf[i-1] != '\r' {
						// discard
						indiceStart = i + 1
						progress = true
						break
					} else {
						// Convert this into a struct

						fmt.Printf("%s\n", buf[indiceStart:i-1])

						// Unmarshal into an object
						respMsg := ResponseMessage{}
						err = json.Unmarshal(buf[indiceStart:i-1], &respMsg)
						if err != nil {
							fmt.Printf("ERROR: %T - %+[1]v\n", err)
						} else {
							// Push the newly created object down the channel
							respMsgChan <- respMsg
						}

						indiceStart = i + 1
						progress = true
						break
					}
				}
			}
		}

		// Shift content to the beginning of the buffer
		if indiceStart != 0 {
			if indiceStop == indiceStart {
				indiceStart = 0
				indiceStop = 0
			} else {
				copy(buf[:], buf[indiceStart:indiceStop])
				indiceStop -= indiceStart
				indiceStart = 0
			}
		}
	}
}

func (esaclient *ESAClient) writer(ReqMsgChan <-chan RequestMessage, stopInformChan chan<- bool) {
	fmt.Println("starting writer goroutine")
	defer fmt.Println("exiting writer goroutine")

	for {
		select {
		case reqMsg, ok := <-ReqMsgChan:
			if !ok {
				close(stopInformChan)
				return
			}

			// Marhsal Request
			bytes, err := json.Marshal(reqMsg)
			if err != nil {
				fmt.Printf("ERROR: %T - %+[1]v\n", err)
				continue
			}

			// esaclient.conn.SetWriteDeadline(time.Now().Add(timeoutDuration))

			// Call Write with timeout
			n, err := esaclient.conn.Write(bytes)
			log.Log(globals.Logger, log.INFO, fmt.Sprintf("write %d bytes from connection", n), nil)

			if err1, ok := err.(*net.OpError); ok {
				// If timeout, continue
				if err1.Timeout() {
					log.Log(globals.Logger, log.INFO, fmt.Sprintf("Timeout: %+v", err1), nil)
					continue
				}
				fmt.Printf("ERROR: %T - %+[1]v\n", err)
				// Before continue, prob need to do something with the potential stuff that is in the buffer!
				// TODO: get the extra data read and update indices!
				continue
			} else if err != nil {
				// If EOF, connection was closed!!
				fmt.Printf("ERROR: %T - %+[1]v\n", err)
				// Before continue, prob need to do something with the potential stuff that is in the buffer!
				// TODO: get the extra data read and update indices!
				continue
			}
		}
	}
}

func (esaclient *ESAClient) Authenticate() error {
	replyChan := make(chan ResponseMessage)
	am := AuthenticationMessage{AppKey: esaclient.appKey, SessionToken: esaclient.sessionToken}
	reqMsg := RequestMessage{Op: "authenticate", AuthenticationMessage: &am}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		// treat response!
		fmt.Printf("%+v\n", resp)
	case <-time.After(3 * time.Second):
		// call timed out
		return errors.New("timeout before getting response")
	}

	return nil
}

func (esaclient *ESAClient) Heartbeat() error {
	replyChan := make(chan ResponseMessage)
	reqMsg := RequestMessage{Op: "heartbeat"}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		// treat response!
		fmt.Printf("%+v\n", resp)
	case <-time.After(3 * time.Second):
		// call timed out
		return errors.New("timeout before getting response")
	}

	return nil
}

func (esaclient *ESAClient) MarketSubscribe(msm MarketSubscriptionMessage) error {
	replyChan := make(chan ResponseMessage)
	reqMsg := RequestMessage{Op: "marketSubscription", MarketSubscriptionMessage: &msm}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		// treat response!
		fmt.Printf("%+v\n", resp)
	case <-time.After(3 * time.Second):
		// call timed out
		return errors.New("timeout before getting response")
	}

	return nil
}

func (esaclient *ESAClient) OrderSubscribe(osm OrderSubscriptionMessage) error {
	replyChan := make(chan ResponseMessage)
	reqMsg := RequestMessage{Op: "orderSubscription", OrderSubscriptionMessage: &osm}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		// treat response!
		fmt.Printf("%+v\n", resp)
	case <-time.After(3 * time.Second):
		// call timed out
		return errors.New("timeout before getting response")
	}

	return nil
}

func (esaclient *ESAClient) getNewID() uint32 {
	return atomic.AddUint32(&esaclient.msgID, 1)
}

type WorkUnit struct {
	req      RequestMessage
	respChan chan ResponseMessage
}

type MarketChangeM struct {
	ID uint32
	MarketChangeMessage
}

type OrderChangeM struct {
	ID uint32
	OrderChangeMessage
}
