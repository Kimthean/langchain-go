package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Base
	Username string
	Email    string    `gorm:"uniqueIndex"`
	Password string    `json:"-"`
	Sessions []Session `gorm:"foreignKey:UserID;references:ID"`
}

type Session struct {
	Base
	UserID uuid.UUID `gorm:"type:uuid"`
	Token  string
}

func MigrateUserModels(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Session{})
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}
