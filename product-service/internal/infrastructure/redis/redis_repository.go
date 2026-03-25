package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisRepository(client *redis.Client, ctx context.Context) *RedisRepository {
	return &RedisRepository{
		Client: client,
		Ctx:    ctx,
	}
}

func (r *RedisRepository) GetByKey(key string) (string, error) {
	result, err := r.Client.Get(r.Ctx, key).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (r *RedisRepository) SetByKey(key, value string) error {
	err := r.Client.Set(r.Ctx, key, value, 60*60*60*time.Second)
	if err != nil {
		return err.Err()
	}

	return nil
}
