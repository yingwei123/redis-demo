package rclient

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestSetAndGetSession(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "session_key"
	value := "session_value"
	expiration := 10 * time.Second

	// Set the session
	err = client.SetSession(key, value, expiration)
	assert.NoError(t, err)

	// Get the session
	result, err := client.GetSession(key)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}

func TestGetSessionNonExistent(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	// Get a non-existent session
	result, err := client.GetSession("non_existent_key")
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestRefreshSessionExpiration(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "session_key"
	value := "session_value"
	initialExpiration := 5 * time.Second
	extendedExpiration := 20 * time.Second

	// Set the session with initial expiration
	err = client.SetSession(key, value, initialExpiration)
	assert.NoError(t, err)

	// Fast-forward to near expiration
	s.FastForward(4 * time.Second)

	// Refresh the session expiration
	err = client.RefreshSessionExpiration(key, extendedExpiration)
	assert.NoError(t, err)

	// Verify the expiration was extended
	ttl, err := rdb.TTL(context.Background(), key).Result()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, ttl.Seconds(), float64(extendedExpiration.Seconds()-1))
}

func TestRefreshSessionExpirationNonExistent(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	// Try to refresh expiration for a non-existent session
	err = client.RefreshSessionExpiration("non_existent_key", 20*time.Second)
	assert.NoError(t, err) // Should return nil without error
}
