package rclient

import (
	"context"
	"strconv"
	"time"
)

// SetInt sets an integer value in the cache
func (c *RedisClient) SetInt(key string, value int, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.SetNX(ctx, key, strconv.Itoa(value), expiration).Err()
}

// GetInt gets an integer value from the cache
func (c *RedisClient) GetInt(key string) (int, error) {
	val, err := c.GetString(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(val)
}
