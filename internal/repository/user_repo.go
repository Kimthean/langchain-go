package repository

import (
	"github.com/Kimthean/go-chat/internal/database"
	"github.com/Kimthean/go-chat/internal/models"
)

func CreateUser(username, email, password string) (*models.User, error) {
	db := database.DB

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	db := database.DB

	user := &models.User{}

	result := db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
