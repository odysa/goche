package lru

import (
	"errors"

	"github.com/odysa/goche/internal/cache"
	"github.com/odysa/goche/internal/utils"
)

type entry[K comparable, V cache.Value] struct {
	key   K
	value V
}

type Cache[K comparable, V cache.Value] struct {
	list     *utils.DoubleList[entry[K, V]]
	keys     map[K]*utils.ListNode[entry[K, V]]
	capacity cache.Size
}

func NewCache[K comparable, V cache.Value](capacity cache.Size) *Cache[K, V] {
	return &Cache[K, V]{
		list:     utils.NewDoubleList[entry[K, V]](),
		keys:     make(map[K]*utils.ListNode[entry[K, V]]),
		capacity: capacity,
	}
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	if node, ok := c.keys[key]; ok {
		c.list.MoveToFront(node)
		return node.Entry().value, nil
	}
	var result V
	return result, errors.New("not found")
}

func (c *Cache[K, V]) Set(key K, value V) error {
	if node, ok := c.keys[key]; ok {
		entry := node.Entry()
		entry.value = value
		c.list.MoveToFront(node)
		return nil
	}

	node := c.list.PushBack(entry[K, V]{key: key, value: value})
	c.keys[key] = node

	if c.reachLimit() {
		c.removeOldest()
	}

	return nil
}

func (c *Cache[K, V]) reachLimit() bool {
	return cache.Size(c.list.Len()) > c.capacity
}

func (c *Cache[K, V]) removeOldest() {
	// it should be unreachable
	if c.list.Len() == 0 {
		return
	}
	oldest := c.list.PopFront()
	delete(c.keys, oldest.Entry().key)
}
