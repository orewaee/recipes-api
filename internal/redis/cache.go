package redis

import (
	"context"
	"errors"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheRepo struct {
	client *redis.Client
}

func NewCacheRepo(ctx context.Context, addr, password string, db int) (repos.CacheRepo, error) {
	options := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}

	client := redis.NewClient(options)

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &CacheRepo{client}, nil
}

func (cache *CacheRepo) Put(ctx context.Context, key, value string, duration time.Duration) error {
	return cache.client.Set(ctx, key, value, duration).Err()
}

func (cache *CacheRepo) Get(ctx context.Context, key string) (string, error) {
	value, err := cache.client.Get(ctx, key).Result()

	if err != nil && errors.Is(err, redis.Nil) {
		return "", domain.ErrNoKey
	}

	if err != nil {
		return "", err
	}

	return value, nil
}

func (cache *CacheRepo) Remove(ctx context.Context, key string) error {
	return cache.client.Del(ctx, key).Err()
}
