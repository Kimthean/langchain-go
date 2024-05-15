package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Kimthean/go-chat/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	log.Println("Connected to PostgreSQL database")

	err = models.MigrateUserModels(DB)
	if err != nil {
		return err
	}

	return nil
}
