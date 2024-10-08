package rclient

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// SetupKeyExpirationNotification enables key expiration notifications in Redis
func (c *RedisClient) SetupKeyExpirationNotification() error {
	ctx := context.Background()

	// Enable keyspace notifications for expired events in Redis
	err := c.Client.ConfigSet(ctx, "notify-keyspace-events", "Ex").Err()
	if err != nil {
		return fmt.Errorf("could not enable keyspace notifications: %v", err)
	}
	return nil
}

// SubscribeToExpiredKeys subscribes to Redis expired key events and returns the PubSub object
func (c *RedisClient) SubscribeToExpiredKeys() (*redis.PubSub, error) {
	// Subscribe to the expired key event notifications
	pubsub := c.Client.Subscribe(context.Background(), "__keyevent@0__:expired")

	// Test if subscription is successful
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not subscribe to expired key events: %v", err)
	}

	return pubsub, nil
}

// ListenForExpiredKeys listens for expired key events and prints the key that expired
func (c *RedisClient) ListenForExpiredKeys(pubsub *redis.PubSub, callback func(string)) {
	ch := pubsub.Channel()

	log.Println("Listening for expired keys...")

	// Loop through the channel and process expired key events
	for msg := range ch {
		fmt.Printf("Key expired: %s\n", msg.Payload)
		callback(msg.Payload)
	}
}
