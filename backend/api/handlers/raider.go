package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hasanaburayyan/raid-bot/backend/db/controllers"
	"github.com/hasanaburayyan/raid-bot/common/models"
)

func GetAllRaiders(c *gin.Context) {
	raiders, err := controllers.GetRaiders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, raiders)
}

func CreateRaider(c *gin.Context) {
	var raider models.Raider
	if err := c.BindJSON(&raider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	raider.ID = uuid.New().String()

	err := controllers.CreateRaider(raider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, raider)
}

func GetRaider(c *gin.Context) {
	id := c.Param("id")

	raider, err := controllers.GetRaiderById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "raider not found"})
		return
	}

	c.JSON(http.StatusAccepted, raider)
}

func UpdateRaider(c *gin.Context) {
	var updatedRaider models.Raider
	if err := c.BindJSON(&updatedRaider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	id := c.Param("id")

	prev, err := controllers.UpdateRaiderById(id, updatedRaider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, prev)
}

func DeleteRaider(c *gin.Context) {
	id := c.Param("id")
	err := controllers.DeleteRaiderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Raider deleted successfully"})
}
