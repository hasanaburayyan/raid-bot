package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckJWT(c *gin.Context) {
	// Get the claims from the context
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found in context"})
		return
	}

	// Ensure the claims are a valid jwt.MapClaims
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims format"})
		return
	}

	// Log the claims for debugging
	log.Println(jwtClaims)

	// Respond with the entire claims map
	c.JSON(http.StatusOK, gin.H{
		"claims": jwtClaims,
	})
}
