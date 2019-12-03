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

type WorkUnit struct {
	req      RequestMessage
	respChan chan ResponseMessage
}

type ConnectionConfig struct {
	// ServerHost is the betfair FQDN/IP
	ServerHost string
	// ServerPort is the betfair server port (443)
	ServerPort uint
	// InsecureSkipVerify sets on/off the TLS cert validation
	InsecureSkipVerify bool

	// ConnectionTimeout specifies the wait time (in milliseconds) before timing out the connection
	ConnectionTimeout int
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

// ESAClient is the client that interacts with betfair Exchange Stream API
// It's thread safe!
type ESAClient struct {
	// Treated as immutable
	appKey       string
	sessionToken string

	// Connection to server
	conn *tls.Conn

	// Atomic
	msgID uint32
	// Connection config
	connConfig atomic.Value
	// ConnectionID - this will be populated after a connection has been established and the first message has been received
	connectionID atomic.Value
	// When sending or receiving messages through channels, how long to wait before giving up (in seconds)
	chanWaitTime uint32
	// Reader buffer size
	readerBufferSize uint32
	// Metrics enable flag (0 - False | 1 - True)
	metricsFlag uint32
	// heartbeat MS interval (bound between 500ms and 5000ms)
	heartbeatMS uint32
	// wait x times the heartbeatMS before sending Heartbeat message (in %) e.g.: 100% means wait for the whole period
	heartbeatMultiplier uint32

	// These channels need to be initialized when the object is created and never changed again!
	// These channels never get closed

	// Gets messages to be sent to the betfair server
	reqMsgChan chan WorkUnit
	// Life cycle management
	// Internal control channels (signal to controller to stop)
	stopChan chan bool
	// Inform back that the controller and its goroutines have successfully shutdown
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

// NewESAClient creates a new esaclient object.
func NewESAClient(appKey string, sessionToken string) ESAClient {
	client := ESAClient{appKey: appKey, sessionToken: sessionToken}
	// Set some defaults
	client.chanWaitTime = 3
	client.readerBufferSize = 8 * 1024 * 1024
	client.heartbeatMS = 1000
	client.heartbeatMultiplier = 150

	client.connectionID.Store("")
	client.connConfig.Store(ConnectionConfig{})

	// Init channels
	client.reqMsgChan = make(chan WorkUnit, 1000)
	client.stopChan = make(chan bool, 1)
	client.stopInformChan = make(chan bool, 1)
	client.heartbeatUpdateChan = make(chan uint32, 10)
	client.authSuccessChan = make(chan bool, 10)
	client.MCMChan = make(chan MarketChangeM, 1000)
	client.OCMChan = make(chan OrderChangeM, 1000)

	return client
}

// ChangeSettings changes some setting on the esaclient.
func (esaclient *ESAClient) ChangeSettings(chanWaitTime uint32, heartbeatMS uint32, heartbeatMultiplier uint32) error {
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

// GetSessionInfo returns the application key, the session token, the connection ID and the current message ID counter.
func (esaclient *ESAClient) GetSessionInfo() (string, string, string, uint32) {
	connID := esaclient.connectionID.Load().(string)

	return esaclient.appKey, esaclient.sessionToken, connID, esaclient.msgID
}

// Connect connects to the betfair server.
// When connected, spawn 4 goroutines: controller, reader, writer and connection tracker.
func (esaclient *ESAClient) Connect(ctx context.Context, connConfig ConnectionConfig) error {

	// If connectionID != "" then there is a connection already!
	if esaclient.connectionID.Load().(string) != "" {
		return ConnectionError{Msg: "connection already established"}
	}

	esaclient.connConfig.Store(connConfig)
	return esaclient.connectionHelper(ctx)
}

// connectionHelper is an helper function used to connect to betfair server
// In case we pass -1 as the retry counter (which means, retry forever) then this function will never return an error
func (esaclient *ESAClient) connectionHelper(ctx context.Context) error {
	connConfig := esaclient.connConfig.Load().(ConnectionConfig)

	config := tls.Config{InsecureSkipVerify: connConfig.InsecureSkipVerify}
	d := net.Dialer{Timeout: time.Duration(connConfig.ConnectionTimeout) * time.Millisecond}
	addr := connConfig.ServerHost + ":" + strconv.Itoa(int(connConfig.ServerPort))

	rpe := NewPolicyExponential(connConfig.Retries, connConfig.MaximumBackoff, 0)

	var cancelled bool
	var wboErr error

	// Retry connection according to retry policy
	for {
		log.Log(globals.Logger, log.INFO, "trying to connect to server", nil)
		err := esaclient.connRetry(ctx, d, addr, &config)

		if err != nil {
			log.Log(globals.Logger, log.ERROR, "failed connecting to server", log.Fields{"attempts": rpe.RetryCount() + 1})
			cancelled, wboErr = rpe.WaitBackOff(ctx)
			if wboErr != nil {
				return fmt.Errorf("failed to connect to server, not trying anymore: %w", err)
			} else if cancelled {
				return fmt.Errorf("request cancelled")
			}
		} else {
			log.Log(globals.Logger, log.INFO, "successfully connected to server", nil)
			return nil
		}
	}
}

func (esaclient *ESAClient) connRetry(ctx context.Context, d net.Dialer, addr string, config *tls.Config) error {
	var err error
	// TODO: Use context if possible
	esaclient.conn, err = tls.DialWithDialer(&d, "tcp", addr, config)
	if err != nil {
		return ConnectionFailedError{Msg: "connecting to betfair failed", Err: err}
	}

	log.Log(globals.Logger, log.INFO, "connection established with server", nil)

	connMsgChan := make(chan ConnectionMessage)

	// Spawn controller in a goroutine
	go esaclient.controller(connMsgChan)

	// Select connMsgChan and Timeout of X seconds
	// If timeout, bring all 4 goroutines down: controller, reader, writer, connection tracker
	select {
	case <-ctx.Done():
		err = esaclient.disconnectHelper()
		return ConnectionFailedError{Msg: "context cancelled while waiting for connection message from betfair", Err: err}
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
	esaclient.stopChan <- true

	// Wait for the stop inform to arrive or after X seconds kill the connection anyway
	select {
	case <-esaclient.stopInformChan:
		// success tearing down
	case <-time.After(time.Duration(esaclient.chanWaitTime) * time.Second):
		// call timed out
		log.Log(globals.Logger, log.ERROR, "timeout while waiting for goroutines to shutdown", nil)
	}

	esaclient.connConfig.Store(ConnectionConfig{})
	esaclient.connectionID.Store("")
	atomic.StoreUint32(&esaclient.msgID, 0)

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

// writer is responsible for sending messages to the betfair server
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

// ========#
// Metrics #
// ========#

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
