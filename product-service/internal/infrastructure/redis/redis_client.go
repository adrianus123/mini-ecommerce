package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_CONNECTION"),
	})
}
