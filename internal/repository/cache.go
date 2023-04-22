package repository

import (
	"context"
	"time"

	"mal-rest-api/internal/config"

	"github.com/redis/go-redis/v9"
)

type CacheRep struct {
	rdb       *redis.Client
	cacheTime time.Duration
}

func NewCache(rdb *redis.Client) *CacheRep {
	return &CacheRep{
		rdb:       rdb,
		cacheTime: config.GetConfig().Redis.CacheTime,
	}
}

func (c *CacheRep) Set(ctx context.Context, key string, value interface{}) error {
	err := c.rdb.Set(ctx, key, value, c.cacheTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheRep) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
