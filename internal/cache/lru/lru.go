package lru

import (
	"container/list"
	"errors"
	"goche/internal/cache"
)

type entry[K comparable, V cache.Value] struct {
	key   K
	value V
}

type Cache[K comparable, V cache.Value] struct {
	list     *list.List
	keys     map[K]*list.Element
	capacity int
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	if node, ok := c.keys[key]; ok {
		c.list.MoveToFront(node)
		return node.Value, nil
	}
	var result V
	return result, errors.New("not found")
}

func (c *Cache[K, V]) Set(key K, value V) (bool, error) {
	if node, ok := c.keys[key]; ok {
		node.Value = value
		c.list.MoveToFront(node)
		return true, nil
	}

	node := c.list.PushFront(value)
	c.keys[key] = node

	if c.list.Len() >= c.capacity {
		c.removeOldest()
	}

	return false, nil
}

func (c *Cache[K, V]) removeOldest() {
	node := c.list.Back()
	if node != nil {
		c.list.Remove(node)

	}
}
