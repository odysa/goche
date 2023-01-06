package lru

import (
	"fmt"
	"testing"

	"github.com/odysa/goche/internal/cache"
	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	lru := NewCache[string, int](10)
	assert.Equal(t, cache.Size(10), lru.capacity)
}

func TestGet(t *testing.T) {
	lru := NewCache[string, int](1)

	err := lru.Set("10", 10)
	assert.Nil(t, err)
	v, err := lru.Get("10")
	assert.Nil(t, err)
	assert.Equal(t, 10, v)

	lru.Set("10", 5)
	v, err = lru.Get("10")
	assert.Nil(t, err)
	assert.Equal(t, 5, v)

	lru.Set("20", 20)
	v, err = lru.Get("10")
	assert.Empty(t, v)
	assert.NotNil(t, err)
}

func TestSet(t *testing.T) {
	lru := NewCache[string, int](1)
	num := 5000
	for i := 0; i < num; i++ {
		lru.Set("10", i)
		v, err := lru.Get("10")
		assert.Nil(t, err)
		assert.Equal(t, i, v)
	}
}

func TestReachLimit(t *testing.T) {
	limit := 41321
	lru := NewCache[string, int](cache.Size(limit))
	for i := 0; i < limit; i++ {
		lru.Set(fmt.Sprint(i), i)
	}
	assert.True(t, lru.reachLimit())
}

func TestRemoveOldest(t *testing.T) {
	limit := cache.Size(5000)
	lru := NewCache[string, int](limit)

	for i := 0; i < int(limit); i++ {
		lru.Set(fmt.Sprint(i), i)
	}

	for i := 0; i < int(limit); i++ {
		v, err := lru.Get(fmt.Sprint(i))
		assert.Nil(t, err)
		assert.Equal(t, i, v)
		lru.removeOldest()
		_, err = lru.Get(fmt.Sprint(i))
		assert.NotNil(t, err)
	}
}
