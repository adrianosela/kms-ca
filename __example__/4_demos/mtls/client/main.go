package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"os"
)

func main() {
	clientPrivFilename := os.Args[1]
	clientCertFilename := os.Args[2]
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

	clientCert, err := tls.LoadX509KeyPair(clientCertFilename, clientPrivFilename)
	if err != nil {
		log.Fatalf("failed to load client tls cert: %v", err)
	}

	conn, err := tls.Dial("tcp", ":2424", &tls.Config{
		ServerName:   "localhost",
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      cp,
		MinVersion:   tls.VersionTLS12,
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
