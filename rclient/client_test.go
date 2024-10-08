package rclient

import (
	"context"
	"testing"

	"redis-demo/config"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// Mock config for Redis client
func mockRedisConfig(s *miniredis.Miniredis) *config.RedisConfig {
	return &config.RedisConfig{
		Host:     s.Host(),
		Port:     s.Port(),
		Password: "", // No password for miniredis
		DB:       0,
	}
}

// Test CreateRedisClient with a valid Redis configuration
func TestCreateRedisClient(t *testing.T) {
	// Start an in-memory Redis server
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	// Mock Redis configuration
	cfg := mockRedisConfig(s)

	// Test creating a Redis client
	client, err := CreateRedisClient(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Check that the client is connected by setting a key in Redis
	err = client.Client.Set(context.Background(), "key", "value", 0).Err()
	assert.NoError(t, err)

	// Retrieve the value to verify the connection
	val, err := client.Client.Get(context.Background(), "key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value", val)
}

// Test CreateRedisClient with an invalid Redis configuration (invalid port)
func TestCreateRedisClientWithInvalidConfig(t *testing.T) {
	// Use an invalid port to simulate a connection error
	cfg := &config.RedisConfig{
		Host: "localhost",
		Port: "99999", // Invalid port
	}

	// Attempt to create the Redis client, expecting an error
	client, err := CreateRedisClient(cfg)
	assert.Error(t, err)
	assert.Nil(t, client)
}

// Test InvalidateCacheKey function
func TestInvalidateCacheKey(t *testing.T) {
	// Start an in-memory Redis server
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	// Mock Redis configuration
	cfg := mockRedisConfig(s)

	// Create Redis client
	client, err := CreateRedisClient(cfg)
	assert.NoError(t, err)

	// Set a key in Redis
	err = client.Client.Set(context.Background(), "key_to_invalidate", "some_value", 0).Err()
	assert.NoError(t, err)

	// Invalidate the cache key
	err = client.InvalidateCacheKey("key_to_invalidate")
	assert.NoError(t, err)

	// Verify that the key no longer exists
	val, err := client.Client.Get(context.Background(), "key_to_invalidate").Result()
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err) // redis.Nil indicates the key doesn't exist
	assert.Empty(t, val)
}

// Test Delete function
func TestDelete(t *testing.T) {
	// Start an in-memory Redis server
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	// Mock Redis configuration
	cfg := mockRedisConfig(s)

	// Create Redis client
	client, err := CreateRedisClient(cfg)
	assert.NoError(t, err)

	// Set a key in Redis
	err = client.Client.Set(context.Background(), "key_to_delete", "some_value", 0).Err()
	assert.NoError(t, err)

	// Delete the key
	err = client.Delete("key_to_delete")
	assert.NoError(t, err)

	// Verify that the key no longer exists
	val, err := client.Client.Get(context.Background(), "key_to_delete").Result()
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err) // redis.Nil indicates the key doesn't exist
	assert.Empty(t, val)
}
