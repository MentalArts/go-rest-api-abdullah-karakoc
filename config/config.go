package config

import (
	"context"
	"fmt"
	"log"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/utils"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global variables for DB and Redis clients
var (
	DB    *gorm.DB
	Redis *redis.Client
)

// ConnectDatabase establishes connections to both PostgreSQL and Redis
func ConnectDatabase() {
	// PostgreSQL connection setup
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

	// Perform database migration
	MigrateDB()

	// Seed Admin User
	SeedAdminUser()

	// Connect to Redis
	ConnectRedis()
}

// MigrateDB runs migrations on the database
func MigrateDB() {
	err := DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{}, &models.User{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
	fmt.Println("Database migrated successfully!")
}

// ConnectRedis connects to the Redis server
func ConnectRedis() {
	// Redis connection parameters from environment variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	// Create a new Redis client
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort), // Redis container address
		Password: "",                                         // No password
		DB:       0,                                          // Default DB
	})

	// Test the Redis connection
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Redis connected successfully!")
}

// SeedAdminUser adds an admin user to the database if not already present
func SeedAdminUser() {
	// Check if the admin already exists
	var admin models.User
	if err := DB.Where("email = ?", "admin@example.com").First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Admin does not exist, create one
			admin = models.User{
				Username: "admin",
				Email:    "admin@gmail.com",
				Password: "adminpassword", // In a real scenario, use a hashed password
				Role:     "admin",         // Set role as admin
			}

			// Hash the password before saving (use a utility function to hash)
			hashedPassword, err := utils.HashPassword(admin.Password)
			if err != nil {
				log.Fatal("Error hashing admin password:", err)
			}

			admin.Password = hashedPassword

			// Save the admin user to the database
			if err := DB.Create(&admin).Error; err != nil {
				log.Fatal("Error creating admin user:", err)
			}
			fmt.Println("Admin user created successfully!")
		} else {
			log.Fatal("Error checking for existing admin user:", err)
		}
	} else {
		fmt.Println("Admin user already exists!")
	}
}
