package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func LoadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open private key file: %w", err)
	}
	defer file.Close()

	keyData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Assert that the key is an RSA private key
	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("key is not an RSA private key")
	}

	return rsaKey, nil
}

func LoadPublicKey(filePath string) (*rsa.PublicKey, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open public key file: %w", err)
	}
	defer file.Close()

	keyData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Assert that the key is an RSA private key
	rsaKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not an RSA private key")
	}

	return rsaKey, nil
}
