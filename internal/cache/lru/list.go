package lru

type Deque[V any] interface {
	PushFront(value V)
	PopFront() DequeNode[V]
	PushBack(value V)
	PopBack() DequeNode[V]
	Size() ListSize
}

type DequeNode[V any] interface {
	GetValue() V
	SetValue(value V)
}

type ListSize = uint64

type DequeList[V any] struct {
	tail, head *DequeListNode[V]
	size       ListSize
}

func (l *DequeList[V]) PushFront(value V) {
	node := newListNode(value, nil, nil)

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}

	l.size += 1
}

func (l *DequeList[V]) PopFront() DequeNode[V] {
	if l.head == nil {
		return nil
	}

	node := l.head
	next := node.next
	node.next = nil

	l.head = next
	l.head.prev = nil
	l.size -= 1

	return node
}

func (l *DequeList[V]) PushBack(value V) {
	node := newListNode(value, nil, nil)

	if l.tail == nil {
		l.tail = node
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}

	l.size += 1
}

func (l *DequeList[V]) PopBack() DequeNode[V] {
	if l.tail == nil {
		return nil
	}
	node := l.tail
	prev := node.prev
	prev.next = nil
	l.tail.prev = nil
	l.tail = prev
	l.size -= 1
	return node
}

func (l *DequeList[V]) Size() ListSize {
	return l.size
}

type DequeListNode[V any] struct {
	value      V
	next, prev *DequeListNode[V]
}

func (l *DequeListNode[V]) GetValue() V {
	return l.value
}

func (l *DequeListNode[V]) SetValue(value V) {
	l.value = value
}

func newListNode[V any](value V, prev, next *DequeListNode[V]) *DequeListNode[V] {
	return &DequeListNode[V]{
		value: value,
		next:  next,
		prev:  prev,
	}
}
