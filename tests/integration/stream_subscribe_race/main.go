package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gustavooferreira/betfair/pkg/auth"
	"github.com/gustavooferreira/betfair/pkg/exchangestream"
)

func main() {
	log.Println("Script starting")

	// var err error

	// AppKey, username, password, certFile, keyFile, connectionTimeout := config()

	// as := auth.NewAuthService(AppKey, username, password, certFile, keyFile, connectionTimeout)
	as := auth.AuthService{AppKey: "app_key", SessionToken: "session_token"}

	// err = as.Login()
	// if err != nil {
	// 	log.Printf("Error while logging in: %s\n", err)
	// }

	// Construct Request
	authMsg := exchangestream.AuthenticationMessage{AppKey: "app_key", SessionToken: "session_token"}
	reqMsg := exchangestream.RequestMessage{Op: "authentication", ID: 100, AuthenticationMessage: &authMsg}
	bytes, err := json.Marshal(reqMsg)
	if err != nil {
		fmt.Printf("ERROR: %T - %+[1]v\n", err)
		return
	}
	fmt.Printf("%s\n", bytes)

	// Deserialize Response
	respMsgString := `{"op":"status", "id":10, "errorMessage":"yolo error message", "errorCode": "UNEXPECTED_ERROR", "connectionId": "123-321", "connectionClosed": true, "statusCode": "SUCCESS"}`

	respMsg := exchangestream.ResponseMessage{}
	err = json.Unmarshal([]byte(respMsgString), &respMsg)
	if err != nil {
		fmt.Printf("ERROR: %T - %+[1]v\n", err)
		return
	}
	fmt.Printf("%+v\n", respMsg)
	fmt.Printf("%+v\n", respMsg.StatusMessage)
	fmt.Printf("Connection closed: %+v\n", *respMsg.StatusMessage.ConnectionClosed)

	if respMsg.StatusMessage.ErrorCode == exchangestream.ErrorCode_UnexpectedError {
		fmt.Println("YOLOOOOO")
	}

	log.Println("Session token: ", as.SessionToken)

	streamStuff(as)

	// err = as.Logout()
	// if err != nil {
	// 	log.Fatalf("Error while logging out: %s\n", err)
	// }

	log.Println("Script ending")
}

func streamStuff(as auth.AuthService) {
	esaclient := exchangestream.NewESAClient(as.AppKey, as.SessionToken)

	fmt.Println("ConnectionID:", esaclient.GetConnectionID(), "|")

	err := esaclient.Connect("127.0.0.1", 8080, true)

	if err != nil {
		log.Println("error while trying to connect")
		log.Println(err.Error())
		return
	}

	fmt.Println("ConnectionID: ", esaclient.GetConnectionID(), "|")

	// err = esaclient.Authenticate(1)
	// if err != nil {
	// 	fmt.Println("error while trying to authenticate")
	// 	return
	// }

	// select {}

	time.Sleep(3 * time.Second)
	log.Println("disconnecting ...")
	err = esaclient.Disconnect()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

// func config() (string, string, string, string, string, uint) {
// 	AppKey, ok := os.LookupEnv("BF_APPKEY")
// 	if !ok {
// 		log.Fatalln("Env var BF_APPKEY missing")
// 	}

// 	username, ok := os.LookupEnv("BF_USERNAME")
// 	if !ok {
// 		log.Fatalln("Env var BF_USERNAME missing")
// 	}

// 	password, ok := os.LookupEnv("BF_PASSWORD")
// 	if !ok {
// 		log.Fatalln("Env var BF_PASSWORD missing")
// 	}

// 	certFile, ok := os.LookupEnv("BF_CERTFILE")
// 	if !ok {
// 		log.Fatalln("Env var BF_CERTFILE missing")
// 	}

// 	keyFile, ok := os.LookupEnv("BF_KEYFILE")
// 	if !ok {
// 		log.Fatalln("Env var BF_KEYFILE missing")
// 	}

// 	var connectionTimeout uint = 10

// 	return AppKey, username, password, certFile, keyFile, connectionTimeout
// }
