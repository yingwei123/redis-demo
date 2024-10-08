package rclient

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// Test SetString and GetString using miniredis (in-memory Redis)
func TestSetGetString(t *testing.T) {
	// Start an in-memory Redis server
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	// Create a Redis client connected to miniredis
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// Create a RedisClient struct wrapping the Redis client
	client := &RedisClient{Client: rdb}

	// Key and value for the test
	key := "test_key"
	value := "test_value"
	expiration := 10 * time.Second

	// Test setting a string value in Redis using SetString
	err = client.SetString(key, value, expiration)
	assert.NoError(t, err)

	// Test getting the string value from Redis using GetString
	returnedValue, err := client.GetString(key)
	assert.NoError(t, err)
	assert.Equal(t, value, returnedValue)

	// Verify that the key has the correct expiration time
	ttl, err := rdb.TTL(context.Background(), key).Result()
	assert.NoError(t, err)
	assert.Greater(t, ttl, 0*time.Second)

	// Try setting the same key again with a different value using SetString
	newValue := "new_test_value"
	err = client.SetString(key, newValue, expiration)
	assert.NoError(t, err)

	// Since SetNX was used, the key should not be overwritten, verify this
	returnedValue, err = client.GetString(key)
	assert.NoError(t, err)
	assert.Equal(t, value, returnedValue) // Value should still be "test_value" (not "new_test_value")
}

// Test GetString for non-existent keys
func TestGetStringNonExistentKey(t *testing.T) {
	// Start an in-memory Redis server
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	// Create a Redis client connected to miniredis
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// Create a RedisClient struct wrapping the Redis client
	client := &RedisClient{Client: rdb}

	// Try to get a non-existent key
	_, err = client.GetString("non_existent_key")
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err) // redis.Nil indicates that the key doesn't exist
}
