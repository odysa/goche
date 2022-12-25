package lru

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
	assert.Equal(t, l.root.Prev().Entry(), "10")
}

func TestPushFront(t *testing.T) {
	l := NewDoubleList[string]()
	l.PushFront("10")
	assert.Equal(t, l.Len(), ListLen(1))
	assert.Equal(t, l.root.Next().Entry(), "10")
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
	assert.Equal(t, n.Entry(), entry)
}
