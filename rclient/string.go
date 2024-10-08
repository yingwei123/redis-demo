package rclient

import (
	"context"
	"time"
)

// SetString sets a raw string value in the cache
func (c *RedisClient) SetString(key string, value string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.SetNX(ctx, key, value, expiration).Err()
}

// GetString gets a raw string value from the cache
func (c *RedisClient) GetString(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.Get(ctx, key).Result()
}
