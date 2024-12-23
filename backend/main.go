package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanaburayyan/raid-bot/backend/api"
)

func main() {
	r := gin.Default()

	api.RegisterRoutes(r)

	r.Run()
}
