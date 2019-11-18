package main

import (
	"crypto/rand"
	"crypto/tls"
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
			err = tlscon.Handshake()
			if err != nil {
				log.Printf("server: TLS handshake: %s", err)
				break
			}
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// defer closeConn(conn)

	msg := "{\"op\":\"connection\",\"connectionId\":\"002-230915140112-174\"}\r\n"
	msgBytes := []byte(msg)

	time.Sleep(4 * time.Second)
	time.Sleep(500 * time.Millisecond)

	n, err := conn.Write(msgBytes[:20])
	log.Printf("Wrote %d bytes - Error: %+v\n", n, err)

	time.Sleep(1500 * time.Millisecond)

	n, err = conn.Write(msgBytes[20:])
	log.Printf("Wrote %d bytes - Error: %+v\n", n, err)

	n, err = conn.Write(msgBytes[:])
	log.Printf("Wrote %d bytes - Error: %+v\n", n, err)

	// n, err := conn.Write([]byte("{\"op\":\"connection\",\"connectionId\":\"002-230915140112-174\"}\r\n{\"op\":\"connect"))
	// log.Printf("Wrote %d bytes - Error: %+v\n", n, err)

	// time.Sleep(1500 * time.Millisecond)

	// n, err = conn.Write([]byte("ion\",\"connectionId\":\"002-230915140112-174\"}\r\n"))
	// log.Printf("Wrote %d bytes - Error: %+v\n", n, err)
}

func closeConn(conn net.Conn) {
	log.Println("server: conn: closed")
	conn.Close()
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
