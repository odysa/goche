package lru

import (
	"errors"
	"goche/internal/cache"
)

type entry[K comparable, V cache.Value] struct {
	key   K
	value V
}

type Cache[K comparable, V cache.Value] struct {
	list     *DoubleList[entry[K, V]]
	keys     map[K]*ListNode[entry[K, V]]
	capacity ListLen
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	if node, ok := c.keys[key]; ok {
		c.list.MoveToFront(node)
		return node.entry.value, nil
	}
	var result V
	return result, errors.New("not found")
}

func (c *Cache[K, V]) Set(key K, value V) (bool, error) {
	if node, ok := c.keys[key]; ok {
		node.entry.value = value
		c.list.MoveToFront(node)
		return true, nil
	}

	node := c.list.PushFront(entry[K, V]{key: key, value: value})
	c.keys[key] = node

	if c.list.Len() >= c.capacity {
		c.removeOldest()
	}

	return false, nil
}

func (c *Cache[K, V]) removeOldest() {
}
