package rclient

import (
	"context"
	"time"
)

// PushToList pushes an element to the left of the list
func (c *RedisClient) PushToList(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LPush(ctx, key, value).Err()
}

// PushToListRight pushes an element to the right of the list
func (c *RedisClient) PushToListRight(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.RPush(ctx, key, value).Err()
}

// PopFromListLeft pops an element from the left of the list
func (c *RedisClient) PopFromListLeft(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LPop(ctx, key).Result()
}

// PopFromListRight pops an element from the right of the list
func (c *RedisClient) PopFromListRight(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.RPop(ctx, key).Result()
}

// GetListRange gets a range of elements from the list
func (c *RedisClient) GetListRange(key string, start, stop int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LRange(ctx, key, start, stop).Result()
}

// GetListLength gets the length of the list
func (c *RedisClient) GetListLength(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LLen(ctx, key).Result()
}

// GetListIndex gets an element from the list by its index
func (c *RedisClient) GetListIndex(key string, index int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LIndex(ctx, key, index).Result()
}

// SetListIndex sets the value of an element in the list by its index
func (c *RedisClient) SetListIndex(key string, index int64, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LSet(ctx, key, index, value).Err()
}

// TrimList trims the list to the specified range (only keep elements between start and stop)
func (c *RedisClient) TrimList(key string, start, stop int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LTrim(ctx, key, start, stop).Err()
}

// RemoveListElements removes elements from the list that match the value
// count > 0: Remove elements moving from head to tail
// count < 0: Remove elements moving from tail to head
// count = 0: Remove all elements equal to the value
func (c *RedisClient) RemoveListElements(key string, count int64, value string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Client.LRem(ctx, key, count, value).Result()
}
