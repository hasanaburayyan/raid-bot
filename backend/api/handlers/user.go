package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/hasanaburayyan/raid-bot/backend/db/controllers"
	"github.com/hasanaburayyan/raid-bot/common/models"
)

func GetAllUsers(c *gin.Context) {
	users, err := controllers.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user.ID = uuid.New().String()

	err := controllers.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var updateUser models.User
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	id := c.Param("id")

	if err := checkOwnership(c, id); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	prev, err := controllers.UpdateUserById(id, updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, prev)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := controllers.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	log.Printf("Found user: %v\n", user)

	c.JSON(http.StatusAccepted, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := checkOwnership(c, id); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err := controllers.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "User deleted successfully"})
}

func checkOwnership(c *gin.Context, requestId string) error {
	claims, ok := c.Get("claims")
	if !ok {
		return fmt.Errorf("no claims set for token")
	}

	claimMap := claims.(jwt.MapClaims)

	claimId, ok := claimMap["id"].(string)
	if !ok {
		return fmt.Errorf("no id claim found")
	}

	role := claimMap["role"].(string)
	if role == "super_admin" {
		return nil
	}

	if claimId != requestId {
		return fmt.Errorf("access denied")
	}

	return nil
}
