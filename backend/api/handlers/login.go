package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hasanaburayyan/raid-bot/backend/auth"
	"github.com/hasanaburayyan/raid-bot/backend/db/controllers"
)

type LoginSuccess struct {
}

func Login(c *gin.Context) {
	var attempt controllers.LoginAttempt

	if err := c.BindJSON(&attempt); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid request body"})
		return
	}

	user, err := controllers.AttemptLogin(attempt)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	}

	token, err := auth.GenerateJWT(auth.PrivateKey, claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	parsedClaims, _ := auth.ParseToken(token, auth.PublicKey)

	c.JSON(http.StatusAccepted, gin.H{
		"token":  token,
		"claims": parsedClaims,
	})
}
