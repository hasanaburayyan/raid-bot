package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hasanaburayyan/raid-bot/backend/api"
	"github.com/hasanaburayyan/raid-bot/backend/auth"
	"github.com/hasanaburayyan/raid-bot/backend/db"
)

func main() {
	auth.LoadKeys()

	r := gin.Default()

	api.RegisterRoutes(r)

	host := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5400")
	dbUser := getEnv("DB_USER", "admin")
	dbPass := getEnv("DB_PASSWORD", "admin")
	dbName := getEnv("DB_NAME", "raidhelper")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, dbUser, dbPass, dbName, dbPort)
	log.Println(dsn)
	db.InitDatabase(dsn)

	db.AutoMigrations()
	db.SeedData()
	r.Run()
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
