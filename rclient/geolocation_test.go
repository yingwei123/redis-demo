package rclient

//this is not testable without running a redis server.
//miniredis is not suitable for this test since it does not geospatial functions.
//redis server version 3.2 is the minimum version that supports geospatial functions.

// import (
// 	"context"
// 	"strings"
// 	"testing"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAddGeoLocationAndGetGeoRadius(t *testing.T) {
// 	// Use a real Redis server for geospatial tests (set your Redis address here)
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	client := &RedisClient{Client: rdb}
// 	key := "geo_test_key"

// 	version, err := rdb.Info(context.Background(), "server").Result()
// 	assert.NoError(t, err)

// 	if !strings.Contains(version, "redis_version:3.2") {
// 		t.Skip("Skipping test: Redis version does not support geospatial commands")
// 	}

// 	// Add geospatial locations
// 	err = client.AddGeoLocation(key, -122.4194, 37.7749, "San Francisco")
// 	assert.NoError(t, err)

// 	err = client.AddGeoLocation(key, -73.935242, 40.730610, "New York")
// 	assert.NoError(t, err)

// 	err = client.AddGeoLocation(key, 2.3522, 48.8566, "Paris")
// 	assert.NoError(t, err)

// 	// Get locations within a 5000 km radius from New York
// 	locations, err := client.GetGeoRadius(key, -73.935242, 40.730610, 5000)
// 	assert.NoError(t, err)

// 	// Check that San Francisco and New York are within the radius
// 	assert.Len(t, locations, 2)
// 	assert.Contains(t, []string{locations[0].Name, locations[1].Name}, "San Francisco")
// 	assert.Contains(t, []string{locations[0].Name, locations[1].Name}, "New York")

// 	// Get locations within a 1000 km radius from Paris
// 	locations, err = client.GetGeoRadius(key, 2.3522, 48.8566, 1000)
// 	assert.NoError(t, err)

// 	// Only Paris should be within the 1000 km radius
// 	assert.Len(t, locations, 1)
// 	assert.Equal(t, "Paris", locations[0].Name)
// }

// func TestGeoRadiusNonExistentKey(t *testing.T) {
// 	// Use a real Redis server for geospatial tests
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	client := &RedisClient{Client: rdb}
// 	key := "non_existent_geo_key"

// 	// Try to get locations from a non-existent key
// 	locations, err := client.GetGeoRadius(key, 0, 0, 100)
// 	assert.NoError(t, err)
// 	assert.Len(t, locations, 0)
// }
