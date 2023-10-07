package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"net"
	"os"
)

func main() {
	privateKeyFilename := os.Args[1]
	csrFilename := os.Args[2]

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate RSA key pair: %v", err)
	}

	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{"Example, Inc."},
			Country:      []string{"US"},
		},
		IPAddresses:        []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		DNSNames:           []string{"localhost"},
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

	// save private key
	privateKeyFile, err := os.Create(privateKeyFilename)
	if err != nil {
		log.Fatalf("failed to open %s for writing key: %v", privateKeyFilename, err)
	}
	defer privateKeyFile.Close()

	privateKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if _, err := privateKeyFile.Write(privateKeyPem); err != nil {
		log.Fatalf("failed to write privkey data to %s: %v", privateKeyFilename, err)
	}

	// save certificate request
	certificateFile, err := os.Create(csrFilename)
	if err != nil {
		log.Fatalf("Failed to open %s for writing csr: %v", csrFilename, err)
	}
	defer certificateFile.Close()

	if _, err := certificateFile.Write(certBytesPem); err != nil {
		log.Fatalf("Failed to write csr data to %s: %v", csrFilename, err)
	}
}
