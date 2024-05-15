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
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Base
	Username string
	Email    string    `gorm:"uniqueIndex"`
	Password string    `json:"-"`
	Sessions []Session `gorm:"foreignKey:UserID"`
}

type Session struct {
	Base
	UserID uuid.UUID `gorm:"type:uuid"`
	Token  string
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func MigrateUserModels(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}
