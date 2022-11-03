package redis_connection

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func InitRedisClient() *redis.Client {
	var (
		opts *redis.Options
	)
	opts = &redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
	}

	client := redis.NewClient(opts)
	return client
}
