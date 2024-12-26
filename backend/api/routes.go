package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanaburayyan/raid-bot/backend/api/handlers"
	"github.com/hasanaburayyan/raid-bot/backend/auth"
)

type RouteGroups struct {
	Public  *gin.RouterGroup
	Private *gin.RouterGroup
	Admin   *gin.RouterGroup
}

func RegisterRoutes(r *gin.Engine) {
	rg := &RouteGroups{
		Public:  r.Group("/"),
		Private: r.Group("/"),
		Admin:   r.Group("/"),
	}

	rg.Private.Use(auth.JWTMiddleware())
	rg.Admin.Use(auth.AdminOnlyMiddleware())

	registerHealthRoutes(rg)
	registerRaiderRoutes(rg)
	registerJWTRoutes(rg)
	registerUserRoutes(rg)
	registerLoginRoutes(rg)
}

func registerHealthRoutes(rg *RouteGroups) {
	rg.Public.GET("/health", handlers.HealthCheck)
}

func registerJWTRoutes(rg *RouteGroups) {
	rg.Private.GET("/jwt", handlers.CheckJWT)
}

func registerRaiderRoutes(rg *RouteGroups) {
	rg.Admin.GET("/raider", handlers.GetAllRaiders)
	rg.Private.GET("/raider/:id", handlers.GetRaider)
	rg.Private.PUT("/raider/:id", handlers.UpdateRaider)
	rg.Private.POST("/raider", handlers.CreateRaider)
	rg.Admin.DELETE("/raider/:id", handlers.DeleteRaider)
}

func registerUserRoutes(rg *RouteGroups) {
	rg.Admin.GET("/user", handlers.GetAllUsers)
	rg.Private.GET("/user/:id", handlers.GetUser)
	rg.Public.POST("/user", handlers.CreateUser)
	rg.Private.PUT("/user/:id", handlers.UpdateUser)
	rg.Admin.DELETE("/user/:id", handlers.DeleteUser)
}

func registerLoginRoutes(rg *RouteGroups) {
	rg.Public.POST("/login", handlers.Login)
}
