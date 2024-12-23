package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanaburayyan/raid-bot/backend/api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	registerHealthRoutes(r)
	registerRaiderRoutes(r)
}

func registerHealthRoutes(r *gin.Engine) {
	r.GET("/health", handlers.HealthCheck)
}

func registerRaiderRoutes(r *gin.Engine) {
	r.GET("/raider", handlers.GetAllRaiders)
	r.GET("/raider/:id", handlers.GetRaider)
	r.PUT("/raider/:id", handlers.UpdateRaider)
	r.POST("/raider", handlers.CreateRaider)
	r.DELETE("/raider/:id", handlers.DeleteRaider)
}
