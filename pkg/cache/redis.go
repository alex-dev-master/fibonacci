package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	Options *redis.Options
	rdb *redis.Client
	ctx context.Context
}

func NewRedis(Host string, Port string, Password string, DB int, ctx context.Context) *Client {
	return &Client{
		Options: &redis.Options{
			Addr:     Host + ":" + Port,
			Password: Password,
			DB: DB,
		},
		ctx: ctx,
	}
}

func (receiver *Client) RunRedis() *Client {
	receiver.rdb = redis.NewClient(receiver.Options)
	return receiver
}

func (receiver *Client) Set(key string, value interface{}, exp time.Duration) error {
	return receiver.rdb.Set(receiver.ctx, key, value, exp).Err()
}

func (receiver *Client) Get(key string) (uint64, error) {
	return receiver.rdb.Get(receiver.ctx, key).Uint64()
}