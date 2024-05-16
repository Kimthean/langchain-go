package repository

import (
	"fmt"

	"github.com/Kimthean/go-chat/internal/database"
	"github.com/Kimthean/go-chat/internal/models"
	"github.com/google/uuid"
)

func CreateUser(username, email, password string) error {
	db := database.DB

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
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

func GetUserByID(id string) (*models.User, error) {
	db := database.DB
	var user models.User

	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func CreateOrUpdateSession(session *models.Session) error {
	db := database.DB

	if session == nil {
		fmt.Println("Session is nil")
		session = &models.Session{}
		fmt.Println("Session is now", session)
	}

	result := db.Where(models.Session{UserID: session.UserID}).Assign(models.Session{Token: session.Token}).FirstOrCreate(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetSessionByToken(token string) (*models.Session, error) {
	db := database.DB

	var session models.Session
	if err := db.Where("token = ?", token).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func GetSessionById(userId uuid.UUID) (*models.Session, error) {
	db := database.DB

	var session models.Session
	if err := db.Where("user_id = ?", userId).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil

}

func DeleteSession(token string) error {
	db := database.DB

	if err := db.Where("token = ?", token).Delete(&models.Session{}).Error; err != nil {
		return err
	}
	return nil
}
