package rclient

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func TestSetJSONAndGetStruct(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_json_key"
	testValue := &TestStruct{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	expiration := 10 * time.Second

	// Set the struct as JSON in Redis
	err = client.SetJSON(key, testValue, expiration)
	assert.NoError(t, err)

	// Get the struct back from Redis
	var result TestStruct
	err = client.GetStruct(key, &result)
	assert.NoError(t, err)

	// Check that the values match
	assert.Equal(t, testValue.Name, result.Name)
	assert.Equal(t, testValue.Age, result.Age)
	assert.Equal(t, testValue.Email, result.Email)
}

func TestGetStructNonExistentKey(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	// Attempt to get a non-existent key
	var result TestStruct
	err = client.GetStruct("non_existent_key", &result)
	assert.Error(t, err)                  // Should return an error for non-existent keys
	assert.Equal(t, TestStruct{}, result) // Expecting the result to be an empty struct
}
