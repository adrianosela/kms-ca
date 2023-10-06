package main

import (
	"context"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/adrianosela/kmsca"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}

	caCertData, err := kmsca.GetSelfSignedCACertificate(
		cfg,
		"alias/my-ca-certificate-key",
		pkix.Name{
			CommonName:   "adrianosela",
			Country:      []string{"CA"},
			Organization: []string{"Adriano Sela Inc."},
		},
		time.Hour*24*365*10, // valid for 10 years
	)
	if err != nil {
		log.Fatalf("failed to get self signed CA certificate: %v", err)
	}

	pemCaCertData := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertData,
	})

	fmt.Println(string(pemCaCertData))
}
