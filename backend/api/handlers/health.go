package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasanaburayyan/raid-bot/common/models"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}

func DummyRaider(c *gin.Context) {
	c.JSON(http.StatusOK, models.Raider{
		ID:        "1",
		DiscordID: "1234567890",
	})
}
