package aping

import (
	"net/http"
	"time"
)

//SetupNetClient sets up a new http client (TLS)
func SetupNetClient(connectionTimeout uint) *http.Client {

	transport := &http.Transport{MaxIdleConns: 200, IdleConnTimeout: 10 * time.Second}
	httpClient := &http.Client{Transport: transport, Timeout: time.Second * time.Duration(connectionTimeout)}
	return httpClient
}

type BetfairAPI struct {
	HttpClient   *http.Client
	AppKey       string
	SessionToken string
}

func NewBetfairAPI(httpClient *http.Client, appKey string, sessionToken string) BetfairAPI {
	bapi := BetfairAPI{HttpClient: httpClient, AppKey: appKey, SessionToken: sessionToken}
	return bapi
}
