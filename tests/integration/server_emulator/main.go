package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Server start")

	certFile, keyFile := config()

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	service := "0.0.0.0:8080"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		tlscon, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlscon.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	msg := `{"op":"connection","connectionId":"002-230915140112-174"}`
	msgBytes := []byte(msg)

	n, err := conn.Write(msgBytes[:20])
	log.Printf("Wrote %d bytes - Error: %+v\n", n, err)

	time.Sleep(3 * time.Second)

	n, err = conn.Write(msgBytes[20:])
	log.Printf("Wrote %d bytes - Error: %+v\n", n, err)

	// buf := make([]byte, 512)
	// for {
	// 	log.Print("server: conn: waiting")
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		if err != nil {
	// 			log.Printf("server: conn: read: %s", err)
	// 		}
	// 		break
	// 	}
	// 	log.Printf("server: conn: echo %q\n", string(buf[:n]))
	// 	n, err = conn.Write(buf[:n])
	// 	log.Printf("server: conn: wrote %d bytes", n)

	// 	if err != nil {
	// 		log.Printf("server: write: %s", err)
	// 		break
	// 	}
	// }
	log.Println("server: conn: closed")
}

func config() (string, string) {
	certFile, ok := os.LookupEnv("SERVER_CERTFILE")
	if !ok {
		log.Fatalln("Env var SERVER_CERTFILE missing")
	}

	keyFile, ok := os.LookupEnv("SERVER_KEYFILE")
	if !ok {
		log.Fatalln("Env var SERVER_KEYFILE missing")
	}

	return certFile, keyFile
}
