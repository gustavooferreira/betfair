package auth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const loginURLGlobal string = "https://identitysso-cert.betfair.com/api/certlogin"
const keepAliveURLGlobal string = "https://identitysso.betfair.com/api/keepAlive"
const logoutURLGlobal string = "https://identitysso.betfair.com/api/logout"

type LoginResponse struct {
	LoginStatus  string `json:"loginStatus"`
	SessionToken string `json:"SessionToken"`
}

type AuthResponse struct {
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
	connectionTimeout uint
}

// NewAuthService creates a AuthService struct.
func NewAuthService(appKey string, username string, password string, certFile string, keyFile string, connectionTimeout uint) AuthService {
	as := AuthService{AppKey: appKey, username: username, password: password, certFile: certFile, keyFile: keyFile, connectionTimeout: connectionTimeout}

	return as
}

// Login authenticates account on the betfair servers and stores valid session token
// to be used on later requests.
func (as *AuthService) Login() (err error) {
	payload := "username=" + as.username + "&password=" + as.password

	// Load client cert
	cert, err := tls.LoadX509KeyPair(as.certFile, as.keyFile)
	if err != nil {
		return err
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{MaxIdleConns: 2, IdleConnTimeout: 10 * time.Second, TLSClientConfig: tlsConfig}
	httpClient := http.Client{Transport: transport, Timeout: time.Second * time.Duration(as.connectionTimeout)}

	req, err := http.NewRequest("POST", loginURLGlobal, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}

	req.Header.Set("X-Application", as.AppKey)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return err
	}

	loginResp := LoginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		return err
	}

	if loginResp.LoginStatus != "SUCCESS" || loginResp.SessionToken == "" {
		return fmt.Errorf("failed login - login status: %s", loginResp.LoginStatus)
	}

	as.SessionToken = loginResp.SessionToken
	return nil
}

// KeepAlive extends the session timeout period.
// At the moment the international (.com) Exchange the current session time is 8 hours.
// If you don't call Keep Alive within the specified timeout period, the session will expire.
// Note:  Session times aren't determined or extended based on API activity.
func (as AuthService) KeepAlive() error {
	if as.SessionToken == "" {
		return errors.New("No session token present")
	}

	httpClient := http.Client{Timeout: time.Second * time.Duration(as.connectionTimeout)}

	req, err := http.NewRequest("GET", keepAliveURLGlobal, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Application", as.AppKey)
	req.Header.Set("X-Authentication", as.SessionToken)
	req.Header.Set("accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return err
	}

	keepaliveResp := AuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(&keepaliveResp)
	if err != nil {
		return err
	}

	if keepaliveResp.Status != "SUCCESS" {
		return fmt.Errorf("failed keep alive request: %s", keepaliveResp.Error)
	}

	return nil
}

// Logout terminates current session.
func (as *AuthService) Logout() error {
	if as.SessionToken == "" {
		return errors.New("No session token present")
	}

	httpClient := http.Client{Timeout: time.Second * time.Duration(as.connectionTimeout)}

	req, err := http.NewRequest("GET", logoutURLGlobal, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Application", as.AppKey)
	req.Header.Set("X-Authentication", as.SessionToken)
	req.Header.Set("accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return err
	}

	logoutResp := AuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(&logoutResp)
	if err != nil {
		return err
	}

	if logoutResp.Status != "SUCCESS" {
		return fmt.Errorf("failed logout request: %s", logoutResp.Error)
	}

	return nil
}
