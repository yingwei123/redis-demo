package rclient

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestSetMapAndGetMap(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_hash"
	mapValue := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
		"field3": "value3",
	}

	// Set the map in Redis
	err = client.SetMap(key, mapValue)
	assert.NoError(t, err)

	// Get the map from Redis
	returnedMap, err := client.GetMap(key)
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		"field1": "value1",
		"field2": "value2",
		"field3": "value3",
	}, returnedMap)
}

func TestGetMapField(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_hash"
	mapValue := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}

	// Set the map in Redis
	err = client.SetMap(key, mapValue)
	assert.NoError(t, err)

	// Get a specific field from the map
	value, err := client.GetMapField(key, "field1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", value)

	// Try to get a non-existent field
	nonExistentValue, err := client.GetMapField(key, "non_existent_field")
	assert.Equal(t, redis.Nil, err)
	assert.Empty(t, nonExistentValue)
}

func TestSetMapAndUpdateField(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_hash"
	mapValue := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}

	// Set the map in Redis
	err = client.SetMap(key, mapValue)
	assert.NoError(t, err)

	// Update one field in the map
	updateField := map[string]interface{}{
		"field2": "new_value2",
	}
	err = client.SetMap(key, updateField)
	assert.NoError(t, err)

	// Get the updated map
	returnedMap, err := client.GetMap(key)
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		"field1": "value1",
		"field2": "new_value2",
	}, returnedMap)
}
