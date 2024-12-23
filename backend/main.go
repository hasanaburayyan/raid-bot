package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hasanaburayyan/raid-bot/backend/api"
	"github.com/hasanaburayyan/raid-bot/backend/db"
)

func main() {
	r := gin.Default()

	api.RegisterRoutes(r)

	dbUrl := os.Getenv("DATABASE_URL")
	conn, err := db.NewConnection(dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	err = db.RunMigrations(conn, "./db/sql")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	r.Run()
}
