package rclient

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

// Set sets a value in the cache with a specified expiration time
func (c *RedisClient) SetJSON(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Marshal the value into JSON
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Set the JSON string in Redis
	err = c.Client.SetNX(ctx, key, jsonValue, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetStruct retrieves a struct from the cache
func (c *RedisClient) GetStruct(key string, dest interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get the value from Redis as a JSON string
	jsonValue, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.New("key does not exist in Redis") // Key does not exist
	} else if err != nil {
		return err
	}

	// Unmarshal the JSON string into the provided destination struct
	err = json.Unmarshal([]byte(jsonValue), dest)
	if err != nil {
		return err
	}

	return nil
}
