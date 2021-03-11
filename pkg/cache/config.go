package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	Options *redis.Options
	ctx context.Context
}

func NewRedisConfig(Host string, Port string, Password string, DB int) *Client {
	return &Client{
		Options: &redis.Options{
			Addr:     Host + ":" + Port,
			Password: Password,
			DB: DB,
		},
	}
}

func (c *Client) RunRedis() (rdb *redis.Client) {
	rdb = redis.NewClient(c.Options)
	return rdb
}