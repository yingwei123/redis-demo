package rclient

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// AddGeoLocation adds a geospatial location to the specified key
func (c *RedisClient) AddGeoLocation(key string, longitude, latitude float64, member string) error {
	if !c.IsGeoSupported() {
		return fmt.Errorf("geospatial commands not supported by this Redis version")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.GeoAdd(ctx, key, &redis.GeoLocation{
		Longitude: longitude,
		Latitude:  latitude,
		Name:      member,
	}).Err()
}

// GetGeoRadius gets members of a geospatial index within a radius
func (c *RedisClient) GetGeoRadius(key string, longitude, latitude, radius float64) ([]redis.GeoLocation, error) {
	if !c.IsGeoSupported() {
		return nil, fmt.Errorf("geospatial commands not supported by this Redis version")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.GeoRadius(ctx, key, longitude, latitude, &redis.GeoRadiusQuery{
		Radius: radius,
		Unit:   "km",
	}).Result()
}

func (c *RedisClient) IsGeoSupported() bool {
	version, err := c.Client.Info(context.Background(), "server").Result()
	if err != nil {
		return false
	}
	return strings.Contains(version, "redis_version:3.2") || strings.Contains(version, "redis_version:4") || strings.Contains(version, "redis_version:5") || strings.Contains(version, "redis_version:6")
}
