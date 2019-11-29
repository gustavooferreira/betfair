// run: go run *.go
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gustavooferreira/betfair/pkg/auth"
	"github.com/gustavooferreira/betfair/pkg/exchangestream"
)

func main() {
	s := fmt.Sprintln("Script starting")
	fmt.Printf(InfoColor, s)

	AppKey, username, password, certFile, keyFile, connectionTimeout, err := config()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	as := auth.NewAuthService(AppKey, username, password, certFile, keyFile, connectionTimeout)

	fmt.Printf(InfoColor, "Logging in ...\n")
	// as.SessionToken = "d4JygGGAl3Hn6/RqGOUhNEFL9y51jZC1euhD8clV3wI="
	err = as.Login()
	if err != nil {
		fmt.Printf("Error while logging in: %s\n", err)
		return
	}

	s = fmt.Sprintln("Session token: ", as.SessionToken)
	fmt.Printf(InfoColor, s)

	time.Sleep(2 * time.Second)

	// globals.Logger = MiniLogger{Level: log.DEBUG}
	streamLogic(as)

	s = fmt.Sprintln("Logging out ...")
	fmt.Printf(InfoColor, s)
	// err = as.Logout()
	// if err != nil {
	// 	fmt.Printf("Error while logging out: %s\n", err)
	// }

	s = fmt.Sprintln("Script ending")
	fmt.Printf(InfoColor, s)
}

func streamLogic(as auth.AuthService) {
	esaclient := exchangestream.NewESAClient(as.AppKey, as.SessionToken)

	err := esaclient.TurnOnMetrics()
	if err != nil {
		s := fmt.Sprintf("ERROR: %+v\n", err)
		fmt.Printf(InfoColor, s)
	} else {
		handler := esaclient.GetMetricsHandler()
		http.Handle("/metrics", handler) //Metrics endpoint for scrapping

		go func() {
			err := http.ListenAndServe(":8080", nil)
			if err != nil {
				s := fmt.Sprintf("ERROR: %+v\n", err)
				fmt.Printf(InfoColor, s)
			}
		}()
	}

	s := fmt.Sprintln("Connecting to betfair server ...")
	fmt.Printf(InfoColor, s)
	err = esaclient.Connect(exchangestream.BetfairHostProd, exchangestream.BetfairPort, false)
	if err != nil {
		var e exchangestream.ConnectionError
		if errors.As(err, &e) {
			s = fmt.Sprintln(err.Error())
			fmt.Printf(InfoColor, s)
		} else {
			s = fmt.Sprintln("Some other error happened while trying to connect to betfair")
			fmt.Printf(InfoColor, s)
			s = fmt.Sprintln(err.Error())
			fmt.Printf(InfoColor, s)
		}

		return
	}

	a, b, c, d := esaclient.GetSessionInfo()
	s = fmt.Sprintf("AppKey: %s | SessionToken: %s | ConnID: %s | MsgID: %d\n", a, b, c, d)
	fmt.Printf(InfoColor, s)

	s = fmt.Sprintln("Authenticating with exchange stream API ...")
	fmt.Printf(InfoColor, s)
	sm, err := esaclient.Authenticate()
	if err != nil {
		s = fmt.Sprintln("error while trying to authenticate")
		fmt.Printf(InfoColor, s)
		// return
	} else {
		s = fmt.Sprintln("exchange stream authentication successful")
		fmt.Printf(InfoColor, s)
		s = fmt.Sprintf("STATUS MESSAGE RESPONSE: %+v\n", sm)
		fmt.Printf(InfoColor, s)
	}

	// Subscribe to markets!
	mf := exchangestream.MarketFilter{CountryCodes: []string{"GB", "ID"}, EventTypeIDs: []string{"7"}}
	mdf := exchangestream.MarketDataFilter{Fields: []exchangestream.PriceData{exchangestream.PriceData_ExBestOffers}}
	msm := exchangestream.MarketSubscriptionMessage{MarketFilter: mf, MarketDataFilter: mdf}

	sm, err = esaclient.MarketSubscribe(msm)
	if err != nil {
		s = fmt.Sprintf("ERROR: %+v\n", err)
		fmt.Printf(InfoColor, s)
	}

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(600 * time.Second)
		c1 <- "result 1"
	}()

	run := true
	for run {
		select {
		case resp := <-esaclient.MCMChan:
			_ = resp
			// s = fmt.Sprintf("%+v\n", resp)
			// fmt.Printf(DebugBoldColor, s)
		case <-c1:
			// stop
			run = false
			break
		}
	}

	s = fmt.Sprintln("Disconnecting from exchange stream API...")
	fmt.Printf(InfoColor, s)
	err = esaclient.Disconnect()
	if err != nil {
		s = fmt.Sprintf("ERROR: %+v\n", err)
		fmt.Printf(InfoColor, s)
	}
}

func config() (string, string, string, string, string, uint, error) {
	AppKey, ok := os.LookupEnv("BF_APPKEY")
	if !ok {
		return "", "", "", "", "", 0, fmt.Errorf("Env var BF_APPKEY missing")
	}

	username, ok := os.LookupEnv("BF_USERNAME")
	if !ok {
		return "", "", "", "", "", 0, fmt.Errorf("Env var BF_USERNAME missing")
	}

	password, ok := os.LookupEnv("BF_PASSWORD")
	if !ok {
		return "", "", "", "", "", 0, fmt.Errorf("Env var BF_PASSWORD missing")
	}

	certFile, ok := os.LookupEnv("BF_CERTFILE")
	if !ok {
		return "", "", "", "", "", 0, fmt.Errorf("Env var BF_CERTFILE missing")
	}

	keyFile, ok := os.LookupEnv("BF_KEYFILE")
	if !ok {
		return "", "", "", "", "", 0, fmt.Errorf("Env var BF_KEYFILE missing")
	}

	var connectionTimeout uint = 10

	return AppKey, username, password, certFile, keyFile, connectionTimeout, nil
}
