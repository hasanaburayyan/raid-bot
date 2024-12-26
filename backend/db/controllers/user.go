package controllers

import (
	"github.com/hasanaburayyan/raid-bot/backend/db"
	"github.com/hasanaburayyan/raid-bot/common/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user models.User) error {
	return db.DB.Create(&user).Error
}

func UpdateUserById(id string, update models.User) (*models.User, error) {
	existing, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	if err := db.DB.Model(&existing).Updates(update).Error; err != nil {
		return nil, err
	}

	return existing, nil
}

func UpdateUserByUsername(username string, update models.User) (*models.User, error) {
	existing, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := db.DB.Model(&existing).Updates(update).Error; err != nil {
		return nil, err
	}

	return existing, nil
}

func DeleteUserById(id string) error {
	usr, err := GetUserById(id)
	if err != nil {
		return err
	}

	if err := db.DB.Delete(&usr).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id string) (*models.User, error) {
	var user models.User

	err := db.DB.First(&user, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := db.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	err := db.DB.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
