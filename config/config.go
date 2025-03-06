package config

import (
	"fmt"
	"log"
	"os"

	"mentalartsapi/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connected successfully!")

	MigrateDB()
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
	fmt.Println("Database migrated successfully!")
}
