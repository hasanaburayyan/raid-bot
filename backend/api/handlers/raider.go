package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hasanaburayyan/raid-bot/backend/db"
	"github.com/hasanaburayyan/raid-bot/common/models"
)

func GetRaiders(c *gin.Context) {
	conn, err := db.NewConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close(context.Background())

	raiders, err := db.GetRaiders(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, raiders)
}

func CreateRaider(c *gin.Context) {
	// Open a database connection
	conn, err := db.NewConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close(context.Background())

	// Bind JSON body to the Raider struct
	var raider models.Raider
	if err := c.BindJSON(&raider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Generate a new UUID for the Raider
	raider.ID = uuid.New().String()

	// Insert the Raider into the database
	err = db.CreateRaider(conn, raider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created Raider
	c.JSON(http.StatusOK, raider)
}
