package kmsca

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

// GetSelfSignedCACertificate returns a self-signed certificate
// signed with a private key managed by the AWS KMS service.
func GetSelfSignedCACertificate(
	cfg aws.Config,
	kmsKeyId string,
	pkixName pkix.Name,
	duration time.Duration,
) ([]byte, error) {
	template := x509.Certificate{
		SerialNumber:          big.NewInt(rand.Int63()), //time seeded serial number
		Subject:               pkixName,
		NotBefore:             time.Now().Add(time.Minute * -5), // valid from 5 minutes ago (allow for clock skews)
		NotAfter:              time.Now().Add(duration),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	kmsSigner, err := NewSigner(cfg, kmsKeyId, types.SigningAlgorithmSpecRsassaPkcs1V15Sha256)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize KMS signer: %v", err)
	}

	certData, err := x509.CreateCertificate(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		&template,
		&template,
		kmsSigner.Public(),
		kmsSigner,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create x509 certificate: %v", err)
	}
	certBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certData,
	})

	return certBytes, nil
}
