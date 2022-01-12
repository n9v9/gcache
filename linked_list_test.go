package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedList(t *testing.T) {
	l := &list[int]{}

	require.Nil(t, l.front, "expected front node on new list to be nil")
	require.Nil(t, l.back, "expected back node on new list to be nil")

	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)

	require.Equal(t, 1, l.popFront())
	require.Equal(t, 2, l.popFront())
	require.Equal(t, 3, l.popFront())

	require.Panics(t, func() { l.popFront() }, "expected popFront on empty list to panic")
}
