package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_Node(t *testing.T) {
	value1 := 13
	node1 := newListNode(value1, nil, nil)
	assert.Equal(t, node1.value, value1)
	assert.Nil(t, node1.next)
	assert.Nil(t, node1.prev)
	value2 := 14
	node2 := newListNode(value2, node1, nil)
	assert.Equal(t, node2.value, value2)
	assert.Equal(t, node2.prev, node1)
	assert.Nil(t, node2.next)
	value3 := 15
	node3 := newListNode(value3, node1, node2)
	assert.Equal(t, node3.value, value3)
	assert.Equal(t, node3.prev, node1)
	assert.Equal(t, node3.next, node2)
}

func TestListNode_GetValue(t *testing.T) {
	value := 256343
	node := newListNode(value, nil, nil)

	assert.Equal(t, node.GetValue(), value)

	strValue := "123"
	strNode := newListNode(strValue, nil, nil)
	assert.Equal(t, strNode.GetValue(), strValue)
}

func TestListNode_SetValue(t *testing.T) {
	value1 := 256343
	node := newListNode(value1, nil, nil)
	assert.Equal(t, node.GetValue(), value1)

	value2 := 6485343
	node.SetValue(value2)

	assert.Equal(t, node.GetValue(), value2)
}
