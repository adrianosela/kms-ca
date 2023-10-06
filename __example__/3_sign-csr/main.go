package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adrianosela/kmsca"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	caCertFilename := os.Args[1]
	csrFilename := os.Args[2]

	caCertPem, err := os.ReadFile(caCertFilename)
	if err != nil {
		log.Fatalf("failed to read CA Certificate PEM: %v", err)
	}
	caCertDer, _ := pem.Decode(caCertPem)

	cert, err := x509.ParseCertificate(caCertDer.Bytes)
	if err != nil {
		log.Fatalf("failed to parse CA Certificate DER: %v", err)
	}

	csrPem, err := os.ReadFile(csrFilename)
	if err != nil {
		log.Fatalf("failed to read CSR PEM: %v", err)
	}
	csrDer, _ := pem.Decode(csrPem)

	csr, err := x509.ParseCertificateRequest(csrDer.Bytes)
	if err != nil {
		log.Fatalf("failed to parse CSR DER: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}

	certdata, err := kmsca.SignCSR(
		cfg,
		"alias/my-ca-certificate-key",
		cert,
		csr,
		time.Hour*24*365, // valid for 1 year
	)
	if err != nil {
		log.Fatalf("failed to sign CSR: %v", err)
	}

	pemCertData := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certdata,
	})

	fmt.Println(string(pemCertData))
}
