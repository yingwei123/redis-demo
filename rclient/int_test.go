package rclient

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestSetIntAndGetInt(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_int_key"
	testValue := 12345
	expiration := 10 * time.Second

	// Set the integer in Redis
	err = client.SetInt(key, testValue, expiration)
	assert.NoError(t, err)

	// Get the integer back from Redis
	result, err := client.GetInt(key)
	assert.NoError(t, err)

	// Check that the values match
	assert.Equal(t, testValue, result)
}

func TestGetIntNonExistentKey(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	// Attempt to get a non-existent key
	result, err := client.GetInt("non_existent_int_key")
	assert.Error(t, err)       // Should return an error since the key doesn't exist
	assert.Equal(t, 0, result) // Result should be 0 for non-existent keys
}

func TestSetIntExpiration(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_int_expiration"
	testValue := 54321
	expiration := 5 * time.Second

	// Set the integer in Redis
	err = client.SetInt(key, testValue, expiration)
	assert.NoError(t, err)

	// Fast forward the time to simulate expiration
	s.FastForward(6 * time.Second)

	// Attempt to get the expired key
	result, err := client.GetInt(key)
	assert.Error(t, err)       // Should return an error since the key has expired
	assert.Equal(t, 0, result) // Result should be 0 since the key expired
}
