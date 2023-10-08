package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	serverPrivFilename := os.Args[1]
	serverCertFilename := os.Args[2]
	trustedCertificate := os.Args[3]

	pemBytes, err := os.ReadFile(trustedCertificate)
	if err != nil {
		log.Fatalf("failed to read trusted cert file: %v", err)
	}
	der, _ := pem.Decode(pemBytes)
	cert, err := x509.ParseCertificate(der.Bytes)
	if err != nil {
		log.Fatalf("failed to parse certificate: %v", err)
	}
	cp := x509.NewCertPool()
	cp.AddCert(cert)

	serverCert, err := tls.LoadX509KeyPair(serverCertFilename, serverPrivFilename)
	if err != nil {
		log.Fatalf("failed to load server tls cert: %v", err)
	}

	listener, err := tls.Listen("tcp", ":2424", &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    cp,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		MinVersion:   tls.VersionTLS12,
	})
	if err != nil {
		log.Fatalf("failed to start TLS listener: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept connection: %v", err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatalf("error reading from connection: %v", err)
		}
	}()

	_, err := io.Copy(conn, os.Stdin)
	if err != nil {
		log.Fatalf("error writing to connection: %v", err)
	}
}
