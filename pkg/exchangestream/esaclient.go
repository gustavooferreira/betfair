package exchangestream

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gustavooferreira/betfair/pkg/globals"
	"github.com/gustavooferreira/betfair/pkg/utils/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
	// Reader buffer
	readerBufferSize uint32
	// Metrics enable (0 - False | 1 - True)
	metricsFlag uint32
	// heartbeat MS interval (bound between 500ms and 5000ms)
	heartbeatMS uint32
	// wait x times the heartbeatMS before sending Heartbeat message (in %) e.g.: 100% means wait for the whole period
	heartbeatMultiplier uint32

	// TODO: These channels need to be initialized when the object is created and never changed again!
	// Never ever ever close these channels!!!!
	reqMsgChan chan WorkUnit

	// Life cycle management
	// Internal channels
	stopChan       chan bool
	stopInformChan chan bool

	// Inform connection tracker that a new update got in
	heartbeatUpdateChan chan uint32
	// Signal auth success
	authSuccessChan chan bool

	// Change Streams
	// Public channel
	MCMChan chan MarketChangeM
	OCMChan chan OrderChangeM

	// Metrics
	readCounter prometheus.Counter
}

func NewESAClient(appKey string, sessionToken string) ESAClient {
	client := ESAClient{appKey: appKey, sessionToken: sessionToken}
	// Set some defaults
	client.connWaitTime = 3
	client.chanWaitTime = 3
	client.readerBufferSize = 8 * 1024 * 1024
	client.heartbeatMS = 1000
	client.heartbeatMultiplier = 150

	client.connectionID.Store("")
	return client
}

func (esaclient *ESAClient) ChangeSettings(connWaitTime uint32, chanWaitTime uint32, heartbeatMS uint32, heartbeatMultiplier uint32) error {
	atomic.StoreUint32(&esaclient.connWaitTime, connWaitTime)
	atomic.StoreUint32(&esaclient.chanWaitTime, chanWaitTime)

	if heartbeatMS < 500 || heartbeatMS > 5000 {
		return fmt.Errorf("heartbeatMS needs to be between 500ms and 5000ms")
	}
	atomic.StoreUint32(&esaclient.heartbeatMS, heartbeatMS)

	if heartbeatMultiplier < 100 {
		return fmt.Errorf("heartbeatMultiplier needs to be at least 100%%")
	}
	atomic.StoreUint32(&esaclient.heartbeatMultiplier, heartbeatMultiplier)

	return nil
}

func (esaclient *ESAClient) GetSessionInfo() (string, string, string, uint32) {
	connID := esaclient.connectionID.Load().(string)

	return esaclient.appKey, esaclient.sessionToken, connID, esaclient.msgID
}

type Config struct {
	// Retries specifies the number of retries allowed.
	// -1 means infinite number of retries
	// 0 means to never retry
	// Any other number greater or equal to 1 means retry X amount of times
	Retries int
	// MaximunBackoff specifies the maximun waiting time between actions (in seconds)
	MaximumBackoff uint
	// Reconnect specifies whether to retry connecting to server upon undesired disconnection
	Reconnect bool
}

// sleepCanBreak is an helper function that sleeps for a specified duration and can be stopped via the context passed in.
func sleepCanBreak(ctx context.Context, sleep time.Duration) (isBreak bool) {
	select {
	case <-ctx.Done():
		isBreak = true
	case <-time.After(sleep):
		isBreak = false
	}
	return
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
		return ConnectionFailedError{Msg: "connecting to betfair failed", Err: err}
	}

	log.Log(globals.Logger, log.INFO, "connection established with server", nil)

	// Init channels
	esaclient.heartbeatUpdateChan = make(chan uint32, 10)
	esaclient.authSuccessChan = make(chan bool, 10)
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
		err = esaclient.disconnectHelper()
		return ConnectionFailedError{Msg: "timeout while waiting for connection message from betfair", Err: err}
	}
}

// Disconnect disconnects from the server
func (esaclient *ESAClient) Disconnect() error {
	// Check that there is a connection to disconnect
	if esaclient.connectionID.Load().(string) == "" {
		return ConnectionError{Msg: "no connection available to disconnect"}
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
		log.Log(globals.Logger, log.ERROR, "timeout while waiting for goroutines to shutdown", nil)
	}

	esaclient.connectionID.Store("")
	atomic.StoreUint32(&esaclient.msgID, 0)

	// Close stream channels
	close(esaclient.MCMChan)
	close(esaclient.OCMChan)

	return esaclient.conn.Close()
}

func (esaclient *ESAClient) controller(connMsgChan chan<- ConnectionMessage) {
	log.Log(globals.Logger, log.INFO, "starting controller goroutine", nil)
	defer log.Log(globals.Logger, log.INFO, "exiting controller goroutine", nil)

	controller(esaclient, connMsgChan)
}

func (esaclient *ESAClient) connTracker(stopChan <-chan bool, stopInformChan chan<- bool) {
	log.Log(globals.Logger, log.INFO, "starting connection tracker goroutine", nil)
	defer log.Log(globals.Logger, log.INFO, "exiting connection tracker goroutine", nil)

	connTracker(esaclient, stopChan, stopInformChan)
}

// reader is responsible for reading all incoming messages and sending the corresponding objects down the channel
func (esaclient *ESAClient) reader(respMsgChan chan<- ResponseMessage, stopChan <-chan bool) {
	log.Log(globals.Logger, log.INFO, "starting reader goroutine", nil)
	defer log.Log(globals.Logger, log.INFO, "exiting reader goroutine", nil)

	reader(esaclient, respMsgChan, stopChan)
}

func (esaclient *ESAClient) writer(reqMsgChan <-chan RequestMessage, stopInformChan chan<- bool) {
	log.Log(globals.Logger, log.INFO, "starting writer goroutine", nil)
	defer log.Log(globals.Logger, log.INFO, "exiting writer goroutine", nil)

	writer(esaclient, reqMsgChan, stopInformChan)
}

// Authenticate authenticates with betfair
func (esaclient *ESAClient) Authenticate() (StatusMessage, error) {
	replyChan := make(chan ResponseMessage)
	am := AuthenticationMessage{AppKey: esaclient.appKey, SessionToken: esaclient.sessionToken}
	reqMsg := RequestMessage{Op: "authentication", AuthenticationMessage: &am}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		if resp.StatusMessage.StatusCode == StatusCode_Success {
			esaclient.authSuccessChan <- true
		}
		return *resp.StatusMessage, nil
	case <-time.After(3 * time.Second):
		return StatusMessage{}, errors.New("timeout before getting response")
	}
}

// Heartbeat sends heartbeat message
func (esaclient *ESAClient) Heartbeat() (StatusMessage, error) {
	replyChan := make(chan ResponseMessage)
	reqMsg := RequestMessage{Op: "heartbeat"}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		return *resp.StatusMessage, nil
	case <-time.After(3 * time.Second):
		return StatusMessage{}, errors.New("timeout before getting response")
	}
}

// MarketSubscribe subscribes to markets
func (esaclient *ESAClient) MarketSubscribe(msm MarketSubscriptionMessage) (StatusMessage, error) {
	replyChan := make(chan ResponseMessage)
	reqMsg := RequestMessage{Op: "marketSubscription", MarketSubscriptionMessage: &msm}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		return *resp.StatusMessage, nil
	case <-time.After(3 * time.Second):
		return StatusMessage{}, errors.New("timeout before getting response")
	}
}

// OrderSubscribe subscribes to orders
func (esaclient *ESAClient) OrderSubscribe(osm OrderSubscriptionMessage) (StatusMessage, error) {
	replyChan := make(chan ResponseMessage)
	reqMsg := RequestMessage{Op: "orderSubscription", OrderSubscriptionMessage: &osm}
	esaclient.reqMsgChan <- WorkUnit{req: reqMsg, respChan: replyChan}

	select {
	case resp := <-replyChan:
		return *resp.StatusMessage, nil
	case <-time.After(3 * time.Second):
		return StatusMessage{}, errors.New("timeout before getting response")
	}
}

func (esaclient *ESAClient) getNewID() uint32 {
	return atomic.AddUint32(&esaclient.msgID, 1)
}

type WorkUnit struct {
	req      RequestMessage
	respChan chan ResponseMessage
}

type MarketChangeM struct {
	ID *uint32
	MarketChangeMessage
}

type OrderChangeM struct {
	ID *uint32
	OrderChangeMessage
}

func (esaclient *ESAClient) TurnOnMetrics() error {
	esaclient.readCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "reader_bytes",
		Help: "Bytes read by the reader",
	})

	//Registering the defined metric with Prometheus
	err := prometheus.Register(esaclient.readCounter)
	if err != nil {
		return err
	}

	atomic.StoreUint32(&esaclient.metricsFlag, 1)
	return nil
}

func (esaclient *ESAClient) TurnOffMetrics() error {
	atomic.StoreUint32(&esaclient.metricsFlag, 0)
	return nil
}

func (esaclient *ESAClient) GetMetricsHandler() http.Handler {
	return promhttp.Handler()
}
