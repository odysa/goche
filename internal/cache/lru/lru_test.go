package lru

import (
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
