package rclient

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// Test SetBit and GetBit using miniredis (in-memory Redis)
func TestSetGetBit(t *testing.T) {
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

	// Key for bit operations
	key := "bitkey"

	// Set a bit at offset 7 to 1 (set the 8th bit to 1)
	err = client.SetBit(key, 7, 1)
	assert.NoError(t, err)

	// Get the bit at offset 7 (should be 1)
	bit, err := client.GetBit(key, 7)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), bit)

	// Set a bit at offset 7 back to 0 (clear the 8th bit)
	err = client.SetBit(key, 7, 0)
	assert.NoError(t, err)

	// Get the bit at offset 7 (should now be 0)
	bit, err = client.GetBit(key, 7)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), bit)

	// Set a bit at offset 0 to 1 (set the 1st bit to 1)
	err = client.SetBit(key, 0, 1)
	assert.NoError(t, err)

	// Get the bit at offset 0 (should be 1)
	bit, err = client.GetBit(key, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), bit)

	// Verify that the key contains the correct bit representation
	val, err := rdb.Get(context.Background(), key).Result()
	assert.NoError(t, err)
	assert.Equal(t, "\x80", val) // '\x80' means the highest bit (8th) is set in a byte.
}
