package rclient

import (
	"context"
	"fmt"
	"redis-demo/config"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient represents a Redis client
type RedisClient struct {
	Client *redis.Client
}

// CreateRedisClient creates a new Redis client and returns a Cache instance
func CreateRedisClient(cfg *config.RedisConfig) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port, // Redis server address
		Password: cfg.Password,              // Password, leave blank if none
		DB:       cfg.DB,                    // Default DB
	})

	// Test the connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	println("Connected to Redis " + cfg.Host + ":" + cfg.Port)

	return &RedisClient{Client: rdb}, nil
}

// Delete deletes a value from the cache
func (c *RedisClient) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

// InvalidateCacheKey deletes the cache key
func (c *RedisClient) InvalidateCacheKey(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the cache key
	err := c.Client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("could not invalidate cache for key %s: %v", key, err)
	}

	fmt.Printf("Cache invalidated for key: %s\n", key)
	return nil
}

// clear all kets from redis
func (c *RedisClient) FlushDB() error {
	err := c.Client.FlushDB(context.Background()).Err()
	if err != nil {
		fmt.Printf("Failed to flush database: %v\n", err)
		return err
	}

	return nil
}
