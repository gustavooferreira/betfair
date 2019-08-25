package exchangestream

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
)

const serverHost string = "stream-api.betfair.com"
const serverPort string = "443"

type Client struct {
	AppKey       string
	SessionToken string
	conn         *tls.Conn
	buffer       *bufio.Scanner
	connectionID string
	idCount      uint32
}

func NewClient(appKey string, sessionToken string) Client {
	client := Client{AppKey: appKey, SessionToken: sessionToken, idCount: 1}
	return client
}

func (esaclient *Client) Connect() error {
	var err error
	esaclient.conn, err = tls.Dial("tcp", serverHost+":"+serverPort, nil)
	if err != nil {
		return err
	}

	esaclient.buffer = bufio.NewScanner(esaclient.conn)

	// FIXME: if it doesn't receive any message it just hangs ....
	_ = esaclient.buffer.Scan()
	message := esaclient.buffer.Text()

	fmt.Println("Message from server: " + message)

	cm := ConnectionMessage{}
	json.Unmarshal([]byte(message), &cm)

	// validate operation is "connection"
	if cm.Op != "connection" {
		return fmt.Errorf("Got OP: %q instead of \"operation\"", cm.Op)
	}

	esaclient.connectionID = cm.ConnectionID

	return nil
}

func (esaclient Client) Disconnect() {
}

func (esaclient *Client) Authenticate() error {

	am := AuthenticationMessage{RequestMessage: RequestMessage{ID: esaclient.getNewID(), Op: "authentication"},
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

func (esaclient *Client) getNewID() uint32 {
	id := esaclient.idCount
	esaclient.idCount++
	return id
}

func (esaclient *Client) Subscribe(marketID string) error {
	sm := SubscriptionMessage{RequestMessage: RequestMessage{Op: "marketSubscription", ID: esaclient.getNewID()},
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

func (esaclient *Client) ReadMessages() (ResponseMessage, error) {

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

func SendHeartbeat() {
	// request = fmt.Sprintf(`{"op": "heartbeat", "id": 2}`)
	// fmt.Println(request)
	// fmt.Fprintf(conn, request+"\r\n")

	// message, _ = buf.ReadString('\n')
	// fmt.Print("Message from server: " + message)
}
