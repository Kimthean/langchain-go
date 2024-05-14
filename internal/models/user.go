package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func MigrateUserModels(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
