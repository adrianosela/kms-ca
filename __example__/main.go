package main

import (
	"context"
	"crypto/x509/pkix"
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

	certBytes, err := kmsca.GetSelfSignedCACertificate(
		cfg,
		"alias/my-ca-certificate-key",
		pkix.Name{
			CommonName:   "adrianosela",
			Country:      []string{"CA"},
			Organization: []string{"Adriano Sela Inc."},
		},
		time.Hour*24*365, // valid for 1 year
	)
	if err != nil {
		log.Fatalf("failed to get self signed CA certificate: %v", err)
	}

	fmt.Println(string(certBytes))
}
