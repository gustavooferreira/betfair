package main

import (
	"log"
	"os"
	"time"

	"github.com/gustavooferreira/betfair/pkg/aping"
	"github.com/gustavooferreira/betfair/pkg/aping/betting"
	"github.com/gustavooferreira/betfair/pkg/auth"
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

	// Get betfair struct
	httpClient := aping.SetupNetClient(connectionTimeout)
	bapi := aping.NewBetfairAPI(httpClient, as.AppKey, as.SessionToken)
	bettingAPI := betting.NewBettingAPI(bapi)

	// List races for the day
	day := time.Now()

	// Tomorrow instead
	day = day.AddDate(0, 0, 1)

	eventYear, eventMonth, eventDay := day.Date()
	from := time.Date(eventYear, eventMonth, eventDay, 0, 0, 0, 0, time.UTC)
	to := time.Date(eventYear, eventMonth, eventDay, 23, 55, 0, 0, time.UTC)
	tr := betting.TimeRange{From: &from, To: &to}

	filter := betting.MarketFilter{EventTypeIds: []string{"7"}, MarketCountries: []string{"GB", "IE"},
		MarketTypeCodes: []string{"WIN"}, MarketStartTime: tr}

	marketProjection := []betting.MarketProjection{betting.MarketProjection_MarketStartTime}
	marketSort := betting.MarketSort_FirstToStart

	marketCatalogue, err := bettingAPI.ListMarketCatalogue(filter, &marketProjection, &marketSort, 100, nil)

	if errB, ok := err.(*betting.BettingAPIError); ok {
		// Just as an example:
		if errB.ErrorCode == betting.APINGExceptionCode_InvalidAppKey {
			log.Fatalln("Invalid App key!")
		}
	} else if err != nil {
		log.Fatalf("Error while getting races for the day: %s\n", err)
	}

	log.Printf("MarketCatalogue: %+v\n", marketCatalogue)

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
