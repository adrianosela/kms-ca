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
	trustedCertificate := os.Args[1]

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

	conn, err := tls.Dial("tcp", ":2424", &tls.Config{
		RootCAs:    cp,
		ServerName: "localhost",
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
