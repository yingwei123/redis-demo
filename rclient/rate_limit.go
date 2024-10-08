package rclient

import (
	"context"
	"time"
)

func (c *RedisClient) RateLimit(key string, limit int, expiration time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := c.Client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		// Set an expiration if this is the first request
		err = c.Client.Expire(ctx, key, expiration).Err()
		if err != nil {
			return false, err
		}
	}

	if count > int64(limit) {
		return false, nil // Rate limit exceeded
	}

	return true, nil // Within rate limit
}
