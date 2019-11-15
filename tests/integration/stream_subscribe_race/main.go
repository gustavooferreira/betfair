package main

import (
	"log"
	"os"

	"github.com/gustavooferreira/betfair/pkg/auth"
	"github.com/gustavooferreira/betfair/pkg/exchangestream"
)

func main() {
	log.Println("Script starting")

	AppKey, username, password, certFile, keyFile, connectionTimeout := config()

	as := auth.NewAuthService(AppKey, username, password, certFile, keyFile, connectionTimeout)

	err := as.Login()
	if err != nil {
		log.Printf("Error while logging in: %s\n", err)

	}

	log.Println("Session token: ", as.SessionToken)

	streamStuff(as)

	err = as.Logout()
	if err != nil {
		log.Fatalf("Error while logging out: %s\n", err)
	}

	log.Println("Script ending")
}

func streamStuff(as auth.AuthService) {
	esaclient := exchangestream.NewESAClient(as.AppKey, as.SessionToken)

	err := esaclient.Connect("127.0.0.1", 8080)

	if err != nil {
		log.Println("error while trying to connect")
		log.Println(err.Error())
		return
	}

	// err = esaclient.Authenticate(1)
	// if err != nil {
	// 	fmt.Println("error while trying to authenticate")
	// 	return
	// }

	log.Println("disconnecting ...")
	esaclient.Disconnect()

}

func config() (string, string, string, string, string, uint) {
	AppKey, ok := os.LookupEnv("BF_APPKEY")
	if !ok {
		log.Fatalln("Env var BF_APPKEY missing")
	}

	username, ok := os.LookupEnv("BF_USERNAME")
	if !ok {
		log.Fatalln("Env var BF_USERNAME missing")
	}

	password, ok := os.LookupEnv("BF_PASSWORD")
	if !ok {
		log.Fatalln("Env var BF_PASSWORD missing")
	}

	certFile, ok := os.LookupEnv("BF_CERTFILE")
	if !ok {
		log.Fatalln("Env var BF_CERTFILE missing")
	}

	keyFile, ok := os.LookupEnv("BF_KEYFILE")
	if !ok {
		log.Fatalln("Env var BF_KEYFILE missing")
	}

	var connectionTimeout uint = 10

	return AppKey, username, password, certFile, keyFile, connectionTimeout
}
