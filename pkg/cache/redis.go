package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisActions interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (uint64, error)
}

type Store struct {
	RedisActions
}

func NewActions(rdb *redis.Client, ctx context.Context) *Store {
	return &Store{
		RedisActions: NewRedisAction(rdb, ctx),
	}
}

