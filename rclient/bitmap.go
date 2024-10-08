package rclient

import (
	"context"
	"time"
)

// SetBit sets or clears the bit at the specified offset
func (c *RedisClient) SetBit(key string, offset int64, value int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.SetBit(ctx, key, offset, value).Err()
}

// GetBit gets the bit value at the specified offset
func (c *RedisClient) GetBit(key string, offset int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.GetBit(ctx, key, offset).Result()
}
