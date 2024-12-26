package auth

import (
	"crypto/rsa"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
)

func GenerateJWT(privateKey *rsa.PrivateKey, claims jwt.MapClaims) (string, error) {
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(privateKey)
}

func LoadKeys() {
	var err error
	PrivateKey, err = LoadPrivateKey("./keys/private.key")
	if err != nil {
		log.Fatalf("error loading private key: %v", err)
	}
	log.Println("Successfully loaded private key")

	PublicKey, err = LoadPublicKey("./keys/public.key")
	if err != nil {
		log.Fatalf("error loading public key: %v", err)
	}
	log.Println("Successfully loaded public key")
}
