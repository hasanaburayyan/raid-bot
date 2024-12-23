package controllers

import (
	"github.com/hasanaburayyan/raid-bot/backend/db"
	"github.com/hasanaburayyan/raid-bot/common/models"
)

func GetRaiderById(id string) (*models.Raider, error) {
	var raider models.Raider

	if err := db.DB.First(&raider, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &raider, nil
}

func CreateRaider(raider models.Raider) error {
	return db.DB.Create(&raider).Error
}

func GetRaiders() ([]models.Raider, error) {
	var raiders []models.Raider

	if err := db.DB.Find(&raiders).Error; err != nil {
		return nil, err
	}

	return raiders, nil
}

func UpdateRaiderById(id string, update models.Raider) (*models.Raider, error) {
	existing, err := GetRaiderById(id)
	if err != nil {
		return nil, err
	}

	if err := db.DB.Model(&existing).Updates(update).Error; err != nil {
		return nil, err
	}

	return existing, nil
}

func DeleteRaiderById(id string) error {
	raider, err := GetRaiderById(id)
	if err != nil {
		return err
	}

	if err := db.DB.Delete(&raider).Error; err != nil {
		return err
	}

	return nil
}
