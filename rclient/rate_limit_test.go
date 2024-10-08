package rclient

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "rate_limit_key"
	limit := 5
	expiration := 10 * time.Second

	// First request, should be allowed
	allowed, err := client.RateLimit(key, limit, expiration)
	assert.NoError(t, err)
	assert.True(t, allowed)

	// Increment the key up to the limit
	for i := 1; i < limit; i++ {
		allowed, err := client.RateLimit(key, limit, expiration)
		assert.NoError(t, err)
		assert.True(t, allowed)
	}

	// Next request should exceed the limit
	allowed, err = client.RateLimit(key, limit, expiration)
	assert.NoError(t, err)
	assert.False(t, allowed)

	// Fast-forward time to expire the key
	s.FastForward(11 * time.Second)

	// After expiration, it should reset and allow the request again
	allowed, err = client.RateLimit(key, limit, expiration)
	assert.NoError(t, err)
	assert.True(t, allowed)
}
