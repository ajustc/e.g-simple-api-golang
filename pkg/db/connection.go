package db

import (
	"api-articles/pkg/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate schema
	err = db.AutoMigrate(&models.Article{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
