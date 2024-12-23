package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanaburayyan/raid-bot/backend/api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", handlers.HealthCheck)
	r.GET("/raider", handlers.DummyRaider)
}
