package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Replace these values with your own configuration
const (
	redisAddr     = "localhost:6379"
	redisPassword = ""
)

var rdb *redis.Client

func InitRedis() {
	// Set up Redis session store
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})
}

func GetRedisClient() *redis.Client {
	return rdb
}

func SetValue(key string, value string) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetValue(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
