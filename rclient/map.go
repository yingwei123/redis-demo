package rclient

import (
	"context"
	"time"
)

// SetMap sets a map value in Redis hash
func (c *RedisClient) SetMap(key string, value map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.HSet(ctx, key, value).Err()
}

// GetMap gets all fields and values of a Redis hash as a map
func (c *RedisClient) GetMap(key string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.HGetAll(ctx, key).Result()
}

// GetMapField gets a specific field of a Redis hash
func (c *RedisClient) GetMapField(key, field string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.HGet(ctx, key, field).Result()
}
