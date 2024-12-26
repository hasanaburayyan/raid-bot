package auth

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Simple Middleware called!")
		c.Next()
	}
}

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeader(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "super_admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "access denied, action can only be done by admin"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Next()
	}
}

func getTokenFromHeader(c *gin.Context) (*jwt.Token, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing Authorization header")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("expired token")
	}

	return token, nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeader(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Next()
	}
}

// ParseToken parses a JWT token string and returns the claims if valid
func ParseToken(tokenString string, publicKey *rsa.PublicKey) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid or expired token: %w", err)
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
