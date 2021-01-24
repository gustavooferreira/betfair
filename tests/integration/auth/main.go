package main

import (
	"log"
	"os"

	"github.com/gustavooferreira/betfair/auth"
)

func main() {
	log.Println("Script starting")

	AppKey, username, password, certFile, keyFile, connectionTimeout := config()

	as := auth.NewAuthService(AppKey, username, password, certFile, keyFile, connectionTimeout)

	err := as.Login()
	if err != nil {
		log.Fatalf("Error while logging in: %s\n", err)
	}

	log.Println("Session token: ", as.SessionToken)

	err = as.KeepAlive()
	if err != nil {
		log.Fatalf("Error while sending keepalive request: %s\n", err)
	}

	err = as.Logout()
	if err != nil {
		log.Fatalf("Error while logging out: %s\n", err)
	}

	log.Println("Script ending")
}

func config() (string, string, string, string, string, uint) {
	AppKey, ok := os.LookupEnv("BF_APPKEY")
	if !ok {
		log.Fatalln("Env var BF_APPKEY missing")
	}

	username, _ := os.LookupEnv("BF_USERNAME")
	if !ok {
		log.Fatalln("Env var BF_USERNAME missing")
	}

	password, _ := os.LookupEnv("BF_PASSWORD")
	if !ok {
		log.Fatalln("Env var BF_PASSWORD missing")
	}

	certFile, _ := os.LookupEnv("BF_CERTFILE")
	if !ok {
		log.Fatalln("Env var BF_CERTFILE missing")
	}

	keyFile, _ := os.LookupEnv("BF_KEYFILE")
	if !ok {
		log.Fatalln("Env var BF_KEYFILE missing")
	}

	var connectionTimeout uint = 10

	return AppKey, username, password, certFile, keyFile, connectionTimeout
}
