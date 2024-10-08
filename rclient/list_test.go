package rclient

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestPushToListAndPopFromList(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_list"

	// Push elements to the left and right of the list
	err = client.PushToList(key, "left1")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "right1")
	assert.NoError(t, err)

	// Pop elements from the left and right
	left, err := client.PopFromListLeft(key)
	assert.NoError(t, err)
	assert.Equal(t, "left1", left)

	right, err := client.PopFromListRight(key)
	assert.NoError(t, err)
	assert.Equal(t, "right1", right)
}

func TestGetListRangeAndLength(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_list"

	// Push elements to the list
	err = client.PushToListRight(key, "elem1")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "elem2")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "elem3")
	assert.NoError(t, err)

	// Get list range
	elements, err := client.GetListRange(key, 0, -1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"elem1", "elem2", "elem3"}, elements)

	// Get list length
	length, err := client.GetListLength(key)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), length)
}

func TestGetListIndexAndSetListIndex(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_list"

	// Push elements to the list
	err = client.PushToListRight(key, "elem1")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "elem2")
	assert.NoError(t, err)

	// Get element by index
	element, err := client.GetListIndex(key, 1)
	assert.NoError(t, err)
	assert.Equal(t, "elem2", element)

	// Set value at index
	err = client.SetListIndex(key, 1, "updated_elem2")
	assert.NoError(t, err)

	// Verify the update
	updatedElement, err := client.GetListIndex(key, 1)
	assert.NoError(t, err)
	assert.Equal(t, "updated_elem2", updatedElement)
}

func TestTrimList(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_list"

	// Push elements to the list
	err = client.PushToListRight(key, "elem1")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "elem2")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "elem3")
	assert.NoError(t, err)

	// Trim the list to keep only the first two elements
	err = client.TrimList(key, 0, 1)
	assert.NoError(t, err)

	// Get the updated list and verify it's been trimmed
	elements, err := client.GetListRange(key, 0, -1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"elem1", "elem2"}, elements)
}

func TestRemoveListElements(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_list"

	// Push elements to the list
	err = client.PushToListRight(key, "elem1")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "elem2")
	assert.NoError(t, err)
	err = client.PushToListRight(key, "elem1")
	assert.NoError(t, err)

	// Remove elements from the list (removing all occurrences of "elem1")
	removedCount, err := client.RemoveListElements(key, 0, "elem1")
	assert.NoError(t, err)
	assert.Equal(t, int64(2), removedCount)

	// Verify the remaining elements
	elements, err := client.GetListRange(key, 0, -1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"elem2"}, elements)
}
