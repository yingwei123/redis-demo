package rclient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (c *RedisClient) Publish(channel string, message interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.Publish(ctx, channel, message).Err()
}

func (c *RedisClient) Subscribe(channel string) *redis.PubSub {
	return c.Client.Subscribe(context.Background(), channel)
}
