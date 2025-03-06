package cache

import (
	"context"
	"fmt"
	"log"

	"mentalartsapi/config"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// SetCache sets a value in the Redis cache with no expiration time (0).
func SetCache(key string, value string) error {
	// Redis'e veri ekleyin
	err := config.Redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Error setting cache: %v", err)
		return err
	}
	fmt.Println("Cache set successfully!")
	return nil
}

// GetCache gets a value from the Redis cache.
func GetCache(key string) (string, error) {
	// Redis'ten veri alın
	val, err := config.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		// Veri bulunamadığında nil döner
		return "", nil
	} else if err != nil {
		// Diğer hatalar
		log.Printf("Error getting cache: %v", err)
		return "", err
	}
	return val, nil
}
