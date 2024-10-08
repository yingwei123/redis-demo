package rclient

//this is not testable without running a redis server since it needs a channel to publish to and subscribe to.
//miniredis is not suitable for this test since it does not support pub/sub.

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/stretchr/testify/assert"
// )

// func TestPublishAndSubscribe(t *testing.T) {
// 	// Use a real Redis server for Pub/Sub testing (set your Redis address here)
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	client := &RedisClient{Client: rdb}
// 	channel := "test_channel"
// 	message := "hello world"

// 	// Subscribe to the channel
// 	sub := client.Subscribe(channel)
// 	defer sub.Close()

// 	// Use a goroutine to receive the message
// 	msgCh := make(chan string)
// 	go func() {
// 		// Wait for the subscription to be ready
// 		_, err := sub.Receive(context.Background())
// 		assert.NoError(t, err)

// 		// Listen for messages
// 		for msg := range sub.Channel() {
// 			msgCh <- msg.Payload
// 		}
// 	}()

// 	// Wait briefly to ensure the subscription is active
// 	time.Sleep(100 * time.Millisecond)

// 	// Publish a message to the channel
// 	err := client.Publish(channel, message)
// 	assert.NoError(t, err)

// 	// Receive the message
// 	receivedMessage := <-msgCh
// 	assert.Equal(t, message, receivedMessage)
// }
