package controllers

import (
	"fmt"

	"github.com/hasanaburayyan/raid-bot/common/models"
)

type LoginAttempt struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AttemptLogin(attempt LoginAttempt) (*models.User, error) {
	user, err := GetUserByUsername(attempt.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid login attempt")
	}

	if user.Password == attempt.Password {
		return user, nil
	}

	return nil, fmt.Errorf("invalid login attempt")
}
