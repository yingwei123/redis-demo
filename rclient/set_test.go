package rclient

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestAddToSetAndGetSetMembers(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_set"
	members := []string{"member1", "member2", "member3"}

	err = client.AddToSet(key, members...)
	assert.NoError(t, err)

	setMembers, err := client.GetSetMembers(key)
	assert.NoError(t, err)
	assert.ElementsMatch(t, members, setMembers)
}

func TestAddToSetWithDuplicateMembers(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	key := "test_set_with_duplicates"
	initialMembers := []string{"member1", "member2"}
	duplicateMembers := []string{"member2", "member3"}

	err = client.AddToSet(key, initialMembers...)
	assert.NoError(t, err)

	err = client.AddToSet(key, duplicateMembers...)
	assert.NoError(t, err)

	expectedMembers := []string{"member1", "member2", "member3"}
	setMembers, err := client.GetSetMembers(key)
	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedMembers, setMembers)
}

func TestGetSetMembersNonExistentKey(t *testing.T) {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{Addr: s.Addr()})
	client := &RedisClient{Client: rdb}

	setMembers, err := client.GetSetMembers("non_existent_set")
	assert.NoError(t, err)
	assert.Empty(t, setMembers)
}
