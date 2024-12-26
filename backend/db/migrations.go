package db

import (
	"log"

	"github.com/hasanaburayyan/raid-bot/common/models"
)

func AutoMigrations() {
	DB.AutoMigrate(
		&models.Raider{},
		&models.Class{},
		&models.User{},
		&models.Role{},
		&models.UserRole{},
		&models.Group{},
		&models.GroupMember{},
	)
}

func SeedData() {
	seedAdminUser()
	seedRoles()
}

func seedRoles() {
	roles := []models.Role{
		{Name: "super_admin"},
		{Name: "user"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := DB.First(&existingRole, "name = ?", role.Name).Error; err == nil {
			log.Printf("Role %s already exists", role.Name)
			continue
		}
		if err := DB.Create(&role).Error; err != nil {
			log.Printf("Failed to create role %s: %v", role.Name, err)
		} else {
			log.Printf("Role %s created", role.Name)
		}
	}
}

func seedAdminUser() {
	var admin models.User
	if err := DB.First(&admin, "username = ?", "admin").Error; err == nil {
		log.Println("Admin user already exists")
		return
	}

	// TODO: use a hashed password
	admin = models.User{
		Username: "admin",
		Password: "admin",
		Role:     "super_admin",
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Fatalf("Error seeding admin user: %v", err)
	}

	log.Println("Admin user created")
}
