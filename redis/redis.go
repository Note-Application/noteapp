package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"

)

// Global Redis client and context
var (
	RedisClient *redis.Client
	RedisCtx    = context.Background() // Global context
)

// InitRedis initializes Redis connection
func InitRedis(host, port string) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   0, // Default DB
	})

	// Test the connection
	_, err := RedisClient.Ping(RedisCtx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")
	return nil
}
