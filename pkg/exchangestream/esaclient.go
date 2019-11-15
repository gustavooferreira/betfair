package exchangestream

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const maxBufCapacity = 1024 * 1024

type ESAClient struct {
	AppKey       string
	SessionToken string
	conn         *tls.Conn

	// RWMutex to access below values
	buffer       *bufio.Scanner
	connectionID string
}

func NewESAClient(appKey string, sessionToken string) ESAClient {
	client := ESAClient{AppKey: appKey, SessionToken: sessionToken}
	return client
}

func (esaclient *ESAClient) Connect(serverHost string, serverPort uint) (err error) {

	// If connectionID != "" then there is a connection already!
	if esaclient.connectionID != "" {
		return fmt.Errorf("connection already established")
	}

	// TODO: REMOVE THIS!
	config := tls.Config{InsecureSkipVerify: true}
	esaclient.conn, err = tls.Dial("tcp", serverHost+":"+strconv.Itoa(int(serverPort)), &config)
	if err != nil {
		return err
	}

	esaclient.buffer = bufio.NewScanner(esaclient.conn)
	buf := make([]byte, maxBufCapacity)
	esaclient.buffer.Buffer(buf, maxBufCapacity)

	// FIXME: if it doesn't receive any message it just hangs ....

	ch := make(chan string)

	go func() {
		res := esaclient.buffer.Scan()
		if !res {
			// ERROR!!!
			// esaclient.buffer.Err().Error()
		}
		message := esaclient.buffer.Text()

		ch <- message
	}()

	select {
	case message := <-ch:
		fmt.Println("Exiting.")
		fmt.Println("Message from server: " + message)

		cm := ConnectionMessage{}
		err = json.Unmarshal([]byte(message), &cm)
		if err != nil {
			// ERROR
		}

		// validate operation is "connection"
		if cm.Op != "connection" {
			return fmt.Errorf("Got OP: %q instead of \"operation\"", cm.Op)
		}

		// WLock!!
		esaclient.connectionID = cm.ConnectionID
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out, exiting.")
	}

	return nil
}

func (esaclient *ESAClient) Disconnect() {
	esaclient.conn.Close()
}

func (esaclient *ESAClient) Authenticate(id uint32) error {

	am := AuthenticationMessage{RequestMessage: RequestMessage{ID: id, Op: "authentication"},
		AppKey:       esaclient.AppKey,
		SessionToken: esaclient.SessionToken}

	amBytes, err := json.Marshal(am)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Message sent to server: %s", amBytes))

	request := string(amBytes) + "\r\n"
	fmt.Fprintf(esaclient.conn, request)

	_ = esaclient.buffer.Scan()
	message := esaclient.buffer.Text()

	fmt.Println("Message from server: " + message)

	rm := ResponseMessage{}
	json.Unmarshal([]byte(message), &rm)

	if rm.Op != "status" {
		return fmt.Errorf("Unexpected message: %s", message)
	}

	sm := rm.StatusMessage

	if sm.StatusCode == "FAILURE" {
		return fmt.Errorf("ErrorCode %s - ErrorMessage %s", sm.ErrorCode, sm.ErrorMessage)
	}

	return nil
}

func (esaclient *ESAClient) Subscribe(marketID string, id uint32) error {
	sm := SubscriptionMessage{RequestMessage: RequestMessage{Op: "marketSubscription", ID: id},
		MarketFilter:     MarketFilter{MarketIDs: []string{marketID}},
		MarketDataFilter: MarketDataFilter{LadderLevels: 3, Fields: []string{"EX_BEST_OFFERS", "EX_MARKET_DEF", "EX_TRADED_VOL", "EX_LTP"}}}

	smBytes, err := json.Marshal(sm)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Message sent to server: %s", smBytes))

	request := string(smBytes) + "\r\n"
	fmt.Fprintf(esaclient.conn, request)

	return nil
}

func (esaclient *ESAClient) ReadMessage() (ResponseMessage, error) {

	// Return object instead of string though!
	if tokenAvail := esaclient.buffer.Scan(); tokenAvail {
		message := esaclient.buffer.Text()

		fmt.Println("Message from server: " + message)

		rm := ResponseMessage{}
		err := json.Unmarshal([]byte(message), &rm)
		if err != nil {
			return rm, err
		}

		return rm, nil
	}

	// *net.OpError
	if err := esaclient.buffer.Err(); err != nil {
		// fmt.Fprintln(os.Stderr, "error:", err)
		// fmt.Fprintf(os.Stderr, "error TYPE: %T\n", err)
		return ResponseMessage{}, err
	}

	return ResponseMessage{}, fmt.Errorf("unknown error")
}

func (esaclient *ESAClient) Heartbeat(id uint32) {
	// request = fmt.Sprintf(`{"op": "heartbeat", "id": 2}`)
	// fmt.Println(request)
	// fmt.Fprintf(conn, request+"\r\n")

	// message, _ = buf.ReadString('\n')
	// fmt.Print("Message from server: " + message)
}
