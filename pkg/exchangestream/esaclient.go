package exchangestream

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

const maxBufCapacity = 1024 * 1024

// timeout of 0.5 seconds
const timeoutDuration = 500 * time.Millisecond

// ESAClient is the client that interacts with betfair Exchange Stream API
// It's thread safe.
type ESAClient struct {
	// Treated as immutable
	appKey       string
	sessionToken string

	// Connection to server
	conn *tls.Conn

	// Atomic
	msgID        uint32
	connectionID atomic.Value

	// Life cycle management
	stopChan       chan bool
	stopInformChan chan bool
}

func NewESAClient(appKey string, sessionToken string) ESAClient {
	client := ESAClient{appKey: appKey, sessionToken: sessionToken}
	client.connectionID.Store("")
	return client
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
		log.Printf("read %d bytes from connection", n)
		if err1, ok := err.(*net.OpError); ok {
			// If timeout, continue
			if err1.Timeout() {
				log.Printf("Timeout: %+v\n", err1)
				continue
			}
			fmt.Printf("ERROR: %T - %[1]+v\n", err)
			// Before continue, prob need to do something with the potential stuff that is in the buffer!
			// TODO: get the extra data read and update indices!
			continue
		} else if err != nil {
			// If EOF, connection was closed!!
			fmt.Printf("ERROR: %T - %[1]+v\n", err)
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
							fmt.Printf("ERROR: %+T - %[1]+v\n", err)
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
				fmt.Printf("ERROR: %T - %[1]+v\n", err)
				continue
			}

			// esaclient.conn.SetWriteDeadline(time.Now().Add(timeoutDuration))

			// Call Write with timeout
			n, err := esaclient.conn.Write(bytes)
			log.Printf("write %d bytes from connection", n)

			if err1, ok := err.(*net.OpError); ok {
				// If timeout, continue
				if err1.Timeout() {
					log.Printf("Timeout: %+v\n", err1)
					continue
				}
				fmt.Printf("ERROR: %+T - %[1]+v\n", err)
				// Before continue, prob need to do something with the potential stuff that is in the buffer!
				// TODO: get the extra data read and update indices!
				continue
			} else if err != nil {
				// If EOF, connection was closed!!
				fmt.Printf("ERROR: %+T - %[1]+v\n", err)
				// Before continue, prob need to do something with the potential stuff that is in the buffer!
				// TODO: get the extra data read and update indices!
				continue
			}
		}
	}
}

func (esaclient *ESAClient) controller(connMsgChan chan<- ConnectionMessage) {
	fmt.Println("starting controller goroutine")
	defer fmt.Println("exiting controller goroutine")

	connPhaseDone := false

	respMsgChan := make(chan ResponseMessage, 1000)
	readerStopChan := make(chan bool)
	reqMsgChan := make(chan RequestMessage, 1000)
	writerStopInformChan := make(chan bool)

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
			}

			fmt.Printf("Message: %+v\n", respMsg)
		}
	}
}

// Connect connects to the server.
// When connected, spawn controller, read and write goroutines!
func (esaclient *ESAClient) Connect(serverHost string, serverPort uint, insecureSkipVerify bool) (err error) {

	// If connectionID != "" then there is a connection already!
	if esaclient.connectionID.Load().(string) != "" {
		return fmt.Errorf("connection already established")
	}

	config := tls.Config{InsecureSkipVerify: insecureSkipVerify}
	d := net.Dialer{Timeout: 3 * time.Second}
	esaclient.conn, err = tls.DialWithDialer(&d, "tcp", serverHost+":"+strconv.Itoa(int(serverPort)), &config)
	if err != nil {
		return err
	}

	fmt.Println("connection established with server")

	esaclient.stopChan = make(chan bool)
	esaclient.stopInformChan = make(chan bool)
	connMsgChan := make(chan ConnectionMessage)

	// Spawn controller in a goroutine
	go esaclient.controller(connMsgChan)

	// Select connMsgChan and Timeout of X seconds
	// If timeout, bring all 3 goroutines down: controller, reader, writer
	select {
	case connMsg := <-connMsgChan:
		esaclient.connectionID.Store(connMsg.ConnectionID)
		return nil

	case <-time.After(3 * time.Second):
		// call timed out
		esaclient.disconnectHelper()

		return errors.New("Timeout while waiting for connection message from betfair")
	}

	return nil
}

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
	case <-time.After(3 * time.Second):
		// call timed out
	}

	esaclient.conn.Close()

	esaclient.connectionID.Store("")

	// Restore id to initial value
	atomic.StoreUint32(&esaclient.msgID, 0)

	return nil
}

func (esaclient *ESAClient) GetConnectionID() string {
	return esaclient.connectionID.Load().(string)
}

// func (esaclient *ESAClient) Authenticate() error {

// 	am := AuthenticationMessage{RequestMessage: RequestMessage{ID: id, Op: "authentication"},
// 		AppKey:       esaclient.appKey,
// 		SessionToken: esaclient.sessionToken}

// 	amBytes, err := json.Marshal(am)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(fmt.Sprintf("Message sent to server: %s", amBytes))

// 	request := string(amBytes) + "\r\n"
// 	fmt.Fprintf(esaclient.conn, request)

// 	_ = esaclient.buffer.Scan()
// 	message := esaclient.buffer.Text()

// 	fmt.Println("Message from server: " + message)

// 	rm := ResponseMessage{}
// 	json.Unmarshal([]byte(message), &rm)

// 	if rm.Op != "status" {
// 		return fmt.Errorf("Unexpected message: %s", message)
// 	}

// 	sm := rm.StatusMessage

// 	if sm.StatusCode == "FAILURE" {
// 		return fmt.Errorf("ErrorCode %s - ErrorMessage %s", sm.ErrorCode, sm.ErrorMessage)
// 	}

// 	return nil
// }

// func (esaclient *ESAClient) Subscribe(marketID string, id uint32) error {
// 	sm := SubscriptionMessage{RequestMessage: RequestMessage{Op: "marketSubscription", ID: id},
// 		MarketFilter:     MarketFilter{MarketIDs: []string{marketID}},
// 		MarketDataFilter: MarketDataFilter{LadderLevels: 3, Fields: []string{"EX_BEST_OFFERS", "EX_MARKET_DEF", "EX_TRADED_VOL", "EX_LTP"}}}

// 	smBytes, err := json.Marshal(sm)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(fmt.Sprintf("Message sent to server: %s", smBytes))

// 	request := string(smBytes) + "\r\n"
// 	fmt.Fprintf(esaclient.conn, request)

// 	return nil
// }

// func (esaclient *ESAClient) ReadMessage() (ResponseMessage, error) {

// 	// Return object instead of string though!
// 	if tokenAvail := esaclient.buffer.Scan(); tokenAvail {
// 		message := esaclient.buffer.Text()

// 		fmt.Println("Message from server: " + message)

// 		rm := ResponseMessage{}
// 		err := json.Unmarshal([]byte(message), &rm)
// 		if err != nil {
// 			return rm, err
// 		}

// 		return rm, nil
// 	}

// 	// *net.OpError
// 	if err := esaclient.buffer.Err(); err != nil {
// 		// fmt.Fprintln(os.Stderr, "error:", err)
// 		// fmt.Fprintf(os.Stderr, "error TYPE: %T\n", err)
// 		return ResponseMessage{}, err
// 	}

// 	return ResponseMessage{}, fmt.Errorf("unknown error")
// }

func (esaclient *ESAClient) Heartbeat() {
	// request = fmt.Sprintf(`{"op": "heartbeat", "id": 2}`)
	// fmt.Println(request)
	// fmt.Fprintf(conn, request+"\r\n")

	// message, _ = buf.ReadString('\n')
	// fmt.Print("Message from server: " + message)
}

func (esaclient *ESAClient) getNewID() uint32 {
	return atomic.AddUint32(&esaclient.msgID, 1)
}
