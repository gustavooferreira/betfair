package exchangestream

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

const maxBufCapacity = 1024 * 1024
const timeoutDuration = 500 * time.Millisecond

type ESAClient struct {
	// Treated as immutable
	appKey       string
	sessionToken string

	// Connection to server
	conn *tls.Conn

	// Atomic
	msgID        uint32
	connectionID atomic.Value
}

func NewESAClient(appKey string, sessionToken string) ESAClient {
	client := ESAClient{appKey: appKey, sessionToken: sessionToken}
	client.connectionID.Store("")
	return client
}

// reader is responsible for reading all incoming messages and sending the corresponding objects down the channel
func (esaclient *ESAClient) reader(respMsgChan chan<- ResponseMessage, stopChan <-chan int) {
	// Create a 8MB buffer for incoming messages
	// var buf []byte = make([]byte, 0, 8*1024*1024)
	bufReader := bufio.NewReader(esaclient.conn)

	for {
		// Call Read with timeout (0.5 seconds)
		esaclient.conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		// Read tokens delimited by newline
		bytes, err := bufReader.ReadBytes('\n')

		// If timeout, continue
		if err1, ok := err.(*net.OpError); ok {
			fmt.Printf("ERROR: %+v\n", err1)

			if err1.Timeout() {
				continue
			}

		} else if err != nil {
			// If EOF, connection was closed!!
			fmt.Printf("ERROR: %+T - %[1]+v\n", err)
			return
		}

		fmt.Printf("%s\n", bytes[:len(bytes)-2])

		cm := ConnectionMessage{}

		// Pass bytes[:len(bytes)-2] to json unmarshelar
		err = json.Unmarshal(bytes[:len(bytes)-2], &cm)
		if err != nil {
			fmt.Printf("ERROR: %+T - %[1]+v\n", err)
		}

		// validate operation is "connection"
		if cm.Op != "connection" {
			// return fmt.Errorf("Got OP: %q instead of \"operation\"", cm.Op)
			fmt.Printf("Got OP: %q instead of \"operation\"", cm.Op)
		}

		esaclient.connectionID.Store(cm.ConnectionID)

		fmt.Println("|", esaclient.connectionID.Load(), "|")

		// check stopChan without blocking, if set then exit function

		// If timeout occurred, then continue (to the beginning of the loop)

		// Check byte slice we got back
		// Iterate through, if contains \n, check if one before was \r then take that len -2 (we don't want \r or \n)

		// Check if buffer is empty! (if not, copy this into that and process it)

		// If it doesn't contain a \n then copy data into our buffer (keep track of the last position occupied)

		// Unmarshal that into an object

		// Push the newly created object down the channel

	}
}

func (esaclient *ESAClient) writer(ReqMsgChan chan RequestMessage, stopInformChan chan<- int) {

}

func (esaclient *ESAClient) controller(stopChan <-chan int, connMsgChan chan<- ConnectionMessage) {
	// Spawn reader and writer goroutines

	respMsgChan := make(chan ResponseMessage, 1000)
	readerStopChan := make(chan int)

	go esaclient.reader(respMsgChan, readerStopChan)

	// If stopChan is closed, trigger reader and writer to end too
	// When getting the connection message, send the data and close the channel
	select {}
}

// Connect connects to the server.
// When connected, spawn read and write goroutines!
func (esaclient *ESAClient) Connect(serverHost string, serverPort uint, insecureSkipVerify bool) (err error) {

	// If connectionID != "" then there is a connection already!
	if esaclient.connectionID.Load().(string) != "" {
		return fmt.Errorf("connection already established")
	}

	config := tls.Config{InsecureSkipVerify: insecureSkipVerify}
	esaclient.conn, err = tls.Dial("tcp", serverHost+":"+strconv.Itoa(int(serverPort)), &config)
	if err != nil {
		return err
	}

	stopChan := make(chan int)
	connMsgChan := make(chan ConnectionMessage)

	// Spawn controller in a goroutine
	go esaclient.controller(stopChan, connMsgChan)

	// Select connMsgChan and Timeout of X seconds
	select {}

	return nil
}

func (esaclient *ESAClient) Disconnect() {
	esaclient.conn.Close()

	// Restore id to initial value
	atomic.StoreUint32(&esaclient.msgID, 0)
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
