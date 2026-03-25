package infrastructure

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRateLimiterRepository struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisRateLimiterRepository(client *redis.Client) *RedisRateLimiterRepository {
	return &RedisRateLimiterRepository{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (r *RedisRateLimiterRepository) Increment(key string, window int) (int64, error) {
	count, err := r.Client.Incr(r.Ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if count == 1 {
		r.Client.Expire(r.Ctx, key, time.Duration(window)*time.Second)
	}

	return count, nil
}
