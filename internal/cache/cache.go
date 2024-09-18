package cache

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisAddr = os.Getenv("REDIS_ADDR")
	password  = os.Getenv("REDIS_PASSWORD")
	Cache     *redis.Client
)

func Connect() {
	redisDBStr := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDBStr)
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       redisDB,
	})

	Cache = client
}

func Set(key string, value interface{}, ttl time.Duration) error {
	ctx := context.Background()
	err := Cache.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string) interface{} {
	ctx := context.Background()
	val, err := Cache.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	return val
}
