package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"mentalartsapi/internal/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func ConnectDatabase() {
	// PostgreSQL bağlantısı
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

	// Veritabanı migrasyonu
	MigrateDB()

	// Redis bağlantısını başlat
	ConnectRedis()
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
	fmt.Println("Database migrated successfully!")
}

func ConnectRedis() {
	// Redis bağlantısı için environment değişkenlerini al
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	// Redis client oluştur
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort), // Redis konteyneri adresi
		Password: "",                                         // Şifre yok
		DB:       0,                                          // Varsayılan DB
	})

	// Redis bağlantısını kontrol et
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Redis connected successfully!")
}
