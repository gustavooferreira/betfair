package auth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const loginURL string = "https://identitysso-cert.betfair.com/api/certlogin"
const keepAliveURL string = "https://identitysso.betfair.com/api/keepAlive"

type LoginResponse struct {
	LoginStatus  string `json:"loginStatus"`
	SessionToken string `json:"SessionToken"`
}

type KeepAliveResponse struct {
	Token   string `json:"token"`
	Product string `json:"product"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

type AuthService struct {
	AppKey            string
	username          string
	password          string
	certFile          string
	keyFile           string
	SessionToken      string
	connectionTimeout int
}

func NewAuthService(appKey string, username string, password string, certFile string, keyFile string) (AuthService, error) {
	bas := AuthService{AppKey: appKey, username: username, password: password, certFile: certFile, keyFile: keyFile, connectionTimeout: 3}

	return bas, nil
}

func (as *AuthService) Login() (sessionToken string, err error) {
	payload := "username=" + as.username + "&password=" + as.password

	// Load client cert
	cert, err := tls.LoadX509KeyPair(as.certFile, as.keyFile)
	if err != nil {
		return sessionToken, err
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{MaxIdleConns: 10, IdleConnTimeout: 10 * time.Second, TLSClientConfig: tlsConfig}
	httpClient := http.Client{Transport: transport, Timeout: time.Second * time.Duration(as.connectionTimeout)}

	req, err := http.NewRequest("POST", loginURL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return sessionToken, err
	}

	req.Header.Set("X-Application", as.AppKey)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)

	if err != nil {
		log.Fatal(fmt.Sprintf("error - Msg: %s", err.Error()))
		return sessionToken, err
	}

	if resp.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("unexpected status code received: %d", resp.StatusCode))
		return sessionToken, err
	}

	loginResp := LoginResponse{}
	json.NewDecoder(resp.Body).Decode(&loginResp)
	resp.Body.Close()

	as.SessionToken = loginResp.SessionToken
	return loginResp.SessionToken, nil
}

func (as AuthService) KeepAlive() error {

	if as.SessionToken == "" {
		return errors.New("No session token present")
	}

	httpClient := http.Client{Timeout: time.Second * time.Duration(as.connectionTimeout)}

	req, err := http.NewRequest("GET", keepAliveURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Application", as.AppKey)
	req.Header.Set("X-Authentication", as.SessionToken)
	req.Header.Set("accept", "application/json")

	resp, err := httpClient.Do(req)

	if err != nil {
		log.Fatal(fmt.Sprintf("error - Msg: %s", err.Error()))
		// return err
	}

	if resp.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("unexpected status code received: %d", resp.StatusCode))
		// return err
	}

	keepaliveResp := KeepAliveResponse{}
	json.NewDecoder(resp.Body).Decode(&keepaliveResp)
	resp.Body.Close()

	fmt.Printf("KeepAlive response: %+v\n", keepaliveResp)

	return nil
}
