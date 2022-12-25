package lru

type ListLen = uint64

type DoubleList[T any] struct {
	root *ListNode[T]
	len  ListLen
}

func NewDoubleList[T any]() *DoubleList[T] {
	var t T
	sentinel := &ListNode[T]{entry: t, prev: nil, next: nil}
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

func (l *DoubleList[T]) PushFront(v T) *ListNode[T] {
	return l.insertAfter(v, l.root)
}

func (l *DoubleList[T]) PushBack(v T) *ListNode[T] {
	return l.insertBefore(v, l.root)
}

func (l *DoubleList[T]) MoveToFront(node *ListNode[T]) {
	// skip the current node
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = l.root
	node.next = l.root.next

	node.prev.next = node
	node.next.prev = node
}

// insert value after node
func (l *DoubleList[T]) insertAfter(v T, node *ListNode[T]) *ListNode[T] {
	newNode := &ListNode[T]{entry: v}

	newNode.prev = node
	newNode.next = node.next

	newNode.prev.next = newNode
	newNode.next.prev = newNode

	l.len += 1
	return newNode
}

func (l *DoubleList[T]) insertBefore(v T, node *ListNode[T]) *ListNode[T] {
	newNode := &ListNode[T]{entry: v}

	newNode.next = node
	newNode.prev = node.prev

	newNode.prev.next = newNode
	newNode.next.prev = newNode

	l.len += 1
	return newNode
}

type ListNode[T any] struct {
	entry      T
	prev, next *ListNode[T]
}

func (n *ListNode[T]) Entry() T {
	return n.entry
}

func (n *ListNode[T]) Prev() *ListNode[T] {
	return n.prev
}

func (n *ListNode[T]) Next() *ListNode[T] {
	return n.next
}
