package lru

type ListLen = uint64

type DoubleList[T any] struct {
	root *ListNode[T]
	len  ListLen
}

func NewDoubleList[T any]() *DoubleList[T] {
	var t T
	sentinel := &ListNode[T]{value: t, prev: nil, next: nil}
	sentinel.prev = sentinel
	sentinel.next = sentinel

	return &DoubleList[T]{
		len:  0,
		root: sentinel,
	}
}

func (l *DoubleList[T]) Len() ListLen {
	return l.len
}

func (l *DoubleList[T]) PushBack(v T) {
}

// insert value after node
func (l *DoubleList[T]) insertAfter(v T, node *ListNode[T]) *ListNode[T] {
	newNode := &ListNode[T]{value: v}

	newNode.prev = node
	newNode.next = node.next

	newNode.prev.next = newNode
	newNode.next.prev = newNode

	l.len += 1
	return newNode
}

func (l *DoubleList[T]) insertBefore(v T, node *ListNode[T]) *ListNode[T] {
	return nil
}

type ListNode[T any] struct {
	value      T
	prev, next *ListNode[T]
}
