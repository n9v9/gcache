package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedList(t *testing.T) {
	l := &List[int]{}

	require.Nil(t, l.front, "expected front node on new list to be nil")
	require.Nil(t, l.back, "expected back node on new list to be nil")

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	require.Equal(t, 1, l.PopFront())
	require.Equal(t, 2, l.PopFront())
	require.Equal(t, 3, l.PopFront())

	require.Panics(t, func() { l.PopFront() }, "expected popFront on empty list to panic")
}
