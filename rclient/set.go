package rclient

import (
	"context"
	"time"
)

// AddToSet adds a member to a Redis set
func (c *RedisClient) AddToSet(key string, members ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.SAdd(ctx, key, members).Err()
}

// GetSetMembers gets all members of a Redis set
func (c *RedisClient) GetSetMembers(key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.SMembers(ctx, key).Result()
}
