package repos

import (
	"context"
	"time"
)

type CacheRepo interface {
	Put(ctx context.Context, key, value string, duration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Remove(ctx context.Context, key string) error
}
