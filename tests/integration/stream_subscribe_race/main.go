package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gustavooferreira/betfair/pkg/auth"
	"github.com/gustavooferreira/betfair/pkg/exchangestream"
	"github.com/gustavooferreira/betfair/pkg/globals"
	"github.com/gustavooferreira/betfair/pkg/utils/log"
)

func main() {
	fmt.Println("Script starting")

	// var err error

	// AppKey, username, password, certFile, keyFile, connectionTimeout := config()

	// as := auth.NewAuthService(AppKey, username, password, certFile, keyFile, connectionTimeout)
	as := auth.AuthService{AppKey: "app_key", SessionToken: "session_token"}

	// err = as.Login()
	// if err != nil {
	// 	log.Printf("Error while logging in: %s\n", err)
	// }

	fmt.Println("Session token: ", as.SessionToken)

	globals.Logger = MiniLogger{Level: log.DEBUG}
	streamStuff(as)

	// err = as.Logout()
	// if err != nil {
	// 	fmt.Fatalf("Error while logging out: %s\n", err)
	// }

	fmt.Println("Script ending")
}

func streamStuff(as auth.AuthService) {
	esaclient := exchangestream.NewESAClient(as.AppKey, as.SessionToken)

	fmt.Print("ConnectionID:")
	fmt.Print(esaclient.GetSessionInfo())
	fmt.Println("|")

	err := esaclient.Connect("127.0.0.1", 8080, true)
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

	fmt.Print("ConnectionID:")
	fmt.Print(esaclient.GetSessionInfo())
	fmt.Println("|")

	// err = esaclient.Authenticate(1)
	// if err != nil {
	// 	fmt.Println("error while trying to authenticate")
	// 	return
	// }

	// select {}

	time.Sleep(3 * time.Second)
	fmt.Println("disconnecting ...")
	err = esaclient.Disconnect()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

// func config() (string, string, string, string, string, uint) {
// 	AppKey, ok := os.LookupEnv("BF_APPKEY")
// 	if !ok {
// 		fmt.Fatalln("Env var BF_APPKEY missing")
// 	}

// 	username, ok := os.LookupEnv("BF_USERNAME")
// 	if !ok {
// 		fmt.Fatalln("Env var BF_USERNAME missing")
// 	}

// 	password, ok := os.LookupEnv("BF_PASSWORD")
// 	if !ok {
// 		fmt.Fatalln("Env var BF_PASSWORD missing")
// 	}

// 	certFile, ok := os.LookupEnv("BF_CERTFILE")
// 	if !ok {
// 		fmt.Fatalln("Env var BF_CERTFILE missing")
// 	}

// 	keyFile, ok := os.LookupEnv("BF_KEYFILE")
// 	if !ok {
// 		fmt.Fatalln("Env var BF_KEYFILE missing")
// 	}

// 	var connectionTimeout uint = 10

// 	return AppKey, username, password, certFile, keyFile, connectionTimeout
// }

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
