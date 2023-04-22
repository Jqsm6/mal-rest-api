package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
}

type Repository struct {
	Cache
}

func NewRepository(ctx context.Context, rdb *redis.Client) *Repository {
	return &Repository{
		Cache: NewCache(rdb),
	}
}
