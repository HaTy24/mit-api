package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Cache *redis.Client
)

func Connect(redisAddr string, password string, redisDB int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       redisDB,
	})
	Cache = client

	return client
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
