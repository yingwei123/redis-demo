package rclient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// SetSession stores a session with the given key, value, and expiration duration
func (c *RedisClient) SetSession(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Store the session in Redis with a TTL
	return c.Client.Set(ctx, key, value, expiration).Err()
}

// GetSession retrieves a session by key
func (c *RedisClient) GetSession(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the session from Redis
	result, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Session not found or expired
		return "", nil
	}
	return result, err
}

// RefreshSessionExpiration extends the expiration of an existing session if it's about to expire
func (c *RedisClient) RefreshSessionExpiration(key string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the session exists
	_, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Session does not exist or is expired
		return nil
	}
	if err != nil {
		return err
	}

	// Refresh the session expiration
	return c.Client.Expire(ctx, key, expiration).Err()
}
