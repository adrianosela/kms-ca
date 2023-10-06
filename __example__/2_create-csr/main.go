package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate RSA key pair: %v", err)
	}

	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   "example.com",
			Organization: []string{"Example, Inc."},
			Country:      []string{"US"},
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	certBytesDer, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		log.Fatalf("Failed to create CSR: %v", err)
	}

	certBytesPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: certBytesDer,
	})

	fmt.Println(string(certBytesPem))
}
