package lru

import "container/list"

type Cache[K comparable, V any] struct {
	keys map[K]*list.Element
	list *list.List
}
