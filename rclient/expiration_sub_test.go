package rclient

//this is not testable without running a redis server.
//miniredis is not suitable for this test since it does not support keyspace notifications.

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/stretchr/testify/assert"
// )

// func TestKeyExpirationNotification(t *testing.T) {
// 	// Use a real Redis server (set your Redis address here)
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	client := &RedisClient{Client: rdb}

// 	// Enable key expiration notifications
// 	err := client.SetupKeyExpirationNotification()
// 	assert.NoError(t, err)

// 	// Subscribe to expired key events
// 	pubsub, err := client.SubscribeToExpiredKeys()
// 	assert.NoError(t, err)
// 	defer pubsub.Close()

// 	// Simulate setting a key with a short expiration time
// 	key := "test_expired_key"
// 	err = client.Client.Set(context.Background(), key, "value", 2*time.Second).Err()
// 	assert.NoError(t, err)

// 	// Use a callback function to capture expired keys
// 	expiredKey := ""
// 	callback := func(key string) {
// 		expiredKey = key
// 	}

// 	// Start a goroutine to listen for expired keys
// 	go client.ListenForExpiredKeys(pubsub, callback)

// 	// Wait for the key to expire
// 	time.Sleep(3 * time.Second)

// 	// Verify that the expired key was detected
// 	assert.Equal(t, key, expiredKey, fmt.Sprintf("Expected key %s to expire, but got %s", key, expiredKey))
// }

// func TestSubscribeToExpiredKeysError(t *testing.T) {
// 	// Use an invalid Redis address to simulate subscription failure
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:9999", // Invalid address
// 	})

// 	client := &RedisClient{Client: rdb}

// 	// Attempt to subscribe to expired key events
// 	_, err := client.SubscribeToExpiredKeys()
// 	assert.Error(t, err, "Expected an error due to invalid Redis address")
// }
