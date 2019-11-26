package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gustavooferreira/betfair/pkg/auth"
	"github.com/gustavooferreira/betfair/pkg/exchangestream"
	"github.com/gustavooferreira/betfair/pkg/globals"
	"github.com/gustavooferreira/betfair/pkg/utils/log"
)

func main() {
	fmt.Println("Script starting")

	// var err error

	AppKey, username, password, certFile, keyFile, connectionTimeout, err := config()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	as := auth.NewAuthService(AppKey, username, password, certFile, keyFile, connectionTimeout)

	fmt.Println("Logging in ...")
	as.SessionToken = "80xwyy64wFT6ogtIBDFjsnb0wDWRBzNplu2AGdqCy5I="
	// err = as.Login()
	// if err != nil {
	// 	fmt.Printf("Error while logging in: %s\n", err)
	// 	return
	// }

	fmt.Println("Session token: ", as.SessionToken)

	globals.Logger = MiniLogger{Level: log.DEBUG}
	streamLogic(as)

	fmt.Println("Logging out ...")
	// err = as.Logout()
	// if err != nil {
	// 	fmt.Printf("Error while logging out: %s\n", err)
	// }

	fmt.Println("Script ending")
}

func streamLogic(as auth.AuthService) {
	esaclient := exchangestream.NewESAClient(as.AppKey, as.SessionToken)

	fmt.Println("Connecting to betfair server ...")
	err := esaclient.Connect(exchangestream.BetfairHostProd, exchangestream.BetfairPort, false)
	if err != nil {
		var e exchangestream.ConnectionError
		if errors.As(err, &e) {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Some other error happened while trying to connect to betfair")
			fmt.Println(err.Error())
		}

		return
	}

	a, b, c, d := esaclient.GetSessionInfo()
	fmt.Printf("AppKey: %s | SessionToken: %s | ConnID: %s | MsgID: %d\n", a, b, c, d)

	// time.Sleep(3 * time.Second)

	fmt.Println("Authenticating with exchange stream API ...")
	err = esaclient.Authenticate()
	if err != nil {
		fmt.Println("error while trying to authenticate")
		// return
	} else {
		fmt.Println("exchange stream authentication successful")
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Disconnecting from exchange stream API...")
	err = esaclient.Disconnect()
	if err != nil {
		fmt.Printf("%+v\n", err)
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

type MiniLogger struct {
	Level log.LogLevel
}

func (ml MiniLogger) Debug(msg string, fields log.Fields) {
	if ml.Level <= log.DEBUG {
		timestamp := time.Now().UTC()
		fmt.Print(generalLogging(msg, "DEBUG", timestamp, fields))
	}
}

func (ml MiniLogger) Info(msg string, fields log.Fields) {
	if ml.Level <= log.INFO {
		timestamp := time.Now().UTC()
		fmt.Print(generalLogging(msg, "INFO", timestamp, fields))
	}
}

func (ml MiniLogger) Warn(msg string, fields log.Fields) {
	if ml.Level <= log.WARN {
		timestamp := time.Now().UTC()
		fmt.Print(generalLogging(msg, "WARN", timestamp, fields))
	}
}

func (ml MiniLogger) Error(msg string, fields log.Fields) {
	if ml.Level <= log.ERROR {
		timestamp := time.Now().UTC()
		fmt.Print(generalLogging(msg, "ERROR", timestamp, fields))
	}
}

func generalLogging(msg string, level string, timestamp time.Time, fields log.Fields) string {
	var container map[string]interface{}

	if len(fields) == 0 {
		container = map[string]interface{}{"message": msg, "level": level, "timestamp": timestamp.Format(time.RFC3339Nano)}
	} else {
		container = map[string]interface{}{"message": msg, "level": level, "extra": fields, "timestamp": timestamp.Format(time.RFC3339Nano)}
	}

	data, err := json.Marshal(container)
	if err != nil {
		return fmt.Sprintf("{\"message\": \"%s\", \"level\":\"%s\", \"timestamp\":\"%s\"}\n", msg, level, timestamp.Format(time.RFC3339Nano))
	}

	return fmt.Sprintln(string(data))
}
