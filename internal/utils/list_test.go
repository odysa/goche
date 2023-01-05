package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDoubleList(t *testing.T) {
	l := NewDoubleList[string]()
	assert.Zero(t, l.Len())
}

func TestPushBack(t *testing.T) {
	l := NewDoubleList[string]()
	l.PushBack("10")
	assert.Equal(t, l.Len(), ListLen(1))
	assert.Equal(t, "10", *l.root.Prev().Entry())
}

func TestPushFront(t *testing.T) {
	l := NewDoubleList[string]()
	l.PushFront("10")
	assert.Equal(t, l.Len(), ListLen(1))
	assert.Equal(t, "10", *l.root.Next().Entry())
}

func TestPopFront(t *testing.T) {
	l := NewDoubleList[string]()
	n1 := l.PushFront("10")
	n2 := l.PushFront("20")

	p1 := l.PopFront()
	assert.Equal(t, l.Len(), ListLen(1))
	assert.Equal(t, p1, n2)
	p2 := l.PopFront()
	assert.Equal(t, l.Len(), ListLen(0))
	assert.Equal(t, p2, n1)
	assert.Nil(t, l.PopFront())
}

func TestPopBack(t *testing.T) {
	l := NewDoubleList[string]()
	n1 := l.PushFront("10")
	n2 := l.PushFront("20")

	p1 := l.PopBack()
	assert.Equal(t, l.Len(), ListLen(1))
	assert.Equal(t, p1, n1)
	p2 := l.PopBack()
	assert.Equal(t, l.Len(), ListLen(0))
	assert.Equal(t, p2, n2)
	assert.Nil(t, l.PopBack())
}

func TestRemove(t *testing.T) {
	l := NewDoubleList[string]()
	n1 := l.PushFront("10")
	n2 := l.PushFront("20")
	l.remove(n1)
	assert.Equal(t, l.Len(), ListLen(1))
	assert.Nil(t, n1.prev)
	assert.Nil(t, n1.next)
	assert.Equal(t, n2.next, l.root)
	assert.Equal(t, n2.prev, l.root)

	l.remove(n2)
	assert.Equal(t, l.Len(), ListLen(0))
	assert.Nil(t, n2.prev)
	assert.Nil(t, n2.next)
	assert.Equal(t, l.root.next, l.root)
	assert.Equal(t, l.root.prev, l.root)
}

func TestListNodePrev(t *testing.T) {
	n1 := &ListNode[string]{
		entry: "10",
		prev:  nil,
		next:  nil,
	}
	assert.Nil(t, n1.Prev())

	n2 := &ListNode[string]{
		entry: "11",
		prev:  nil,
		next:  nil,
	}
	n1.prev = n2

	assert.Equal(t, n1.Prev(), n2)
}
func TestListNodeNext(t *testing.T) {
	n1 := &ListNode[string]{
		entry: "10",
		prev:  nil,
		next:  nil,
	}
	assert.Nil(t, n1.Next())

	n2 := &ListNode[string]{
		entry: "11",
		prev:  nil,
		next:  nil,
	}
	n1.next = n2

	assert.Equal(t, n1.Next(), n2)
}
func TestListNodeEntry(t *testing.T) {
	entry := "12321312"
	n := &ListNode[string]{
		entry: entry,
		prev:  nil,
		next:  nil,
	}
	assert.Equal(t, entry, *n.Entry())
}
