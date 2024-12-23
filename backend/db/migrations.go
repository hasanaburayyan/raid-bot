package db

import (
	"github.com/hasanaburayyan/raid-bot/common/models"
)

func AutoMigrations() {
	DB.AutoMigrate(&models.Raider{})
	DB.AutoMigrate(&models.Class{})
}
