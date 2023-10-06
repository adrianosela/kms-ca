package kmsca

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

// Signer can sign arbitrary bytes.
type Signer struct {
	kmsClient   *kms.Client
	kmsKeyId    string
	signingAlgo types.SigningAlgorithmSpec

	publicKey crypto.PublicKey
}

// ensure Signer implements crypto.Signer.
var _ crypto.Signer = (*Signer)(nil)

// NewSigner returns a new signer
func NewSigner(cfg aws.Config, kmsKeyId string, signingAlgo types.SigningAlgorithmSpec) (*Signer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	kmsClient := kms.NewFromConfig(cfg)
	getPublicKeyOutput, err := kmsClient.GetPublicKey(ctx, &kms.GetPublicKeyInput{
		KeyId: aws.String(kmsKeyId),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get public key %s from KMS: %v", kmsKeyId, err)
	}
	pemBlock, _ := pem.Decode(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: getPublicKeyOutput.PublicKey,
	}))
	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key bytes from KMS: %s", err)
	}

	return &Signer{
		kmsClient:   kmsClient,
		kmsKeyId:    kmsKeyId,
		signingAlgo: signingAlgo,
		publicKey:   publicKey,
	}, nil
}

// Public implements the Signer interface.
func (s *Signer) Public() crypto.PublicKey {
	return s.publicKey
}

// Sign delegates the signing operation to KMS.
func (s *Signer) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	ctx := context.Background()

	signOutput, err := s.kmsClient.Sign(ctx, &kms.SignInput{
		KeyId:            aws.String(s.kmsKeyId),
		Message:          digest,
		MessageType:      types.MessageTypeDigest,
		SigningAlgorithm: s.signingAlgo,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sign digest with KMS: %v", err)
	}

	return signOutput.Signature, err
}
