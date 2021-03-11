package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Actions struct {
	rdb redis.Client
	ctx context.Context
}

func NewRedisAction(rdb *redis.Client, ctx context.Context) *Actions {
	return &Actions{
		rdb: *rdb,
		ctx: ctx,
	}
}

func (receiver *Actions) Set(key string, value interface{}, exp time.Duration) error {
	return receiver.rdb.Set(receiver.ctx, key, value, exp).Err()
}

func (receiver *Actions) Get(key string) (uint64, error) {
	return receiver.rdb.Get(receiver.ctx, key).Uint64()
}