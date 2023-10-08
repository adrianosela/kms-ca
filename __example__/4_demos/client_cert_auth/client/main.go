package main

import (
	"crypto/tls"
	"io"
	"log"
	"os"
)

func main() {
	clientPrivFilename := os.Args[1]
	clientCertFilename := os.Args[2]

	clientCert, err := tls.LoadX509KeyPair(clientCertFilename, clientPrivFilename)
	if err != nil {
		log.Fatalf("failed to load client tls cert: %v", err)
	}

	conn, err := tls.Dial("tcp", ":2424", &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true, // do not verify server cert
	})
	if err != nil {
		log.Fatalf("failed to dial TLS: %v", err)
	}

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatalf("error reading from connection: %v", err)
		}
	}()

	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		log.Fatalf("error writing to connection: %v", err)
	}
}
