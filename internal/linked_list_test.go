package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedList(t *testing.T) {
	t.Run("PopFront on empty list panics", func(t *testing.T) {
		l := &List[int]{}
		require.Panics(t, func() { l.PopFront() })
	})

	t.Run("PushBack", func(t *testing.T) {
		l := &List[int]{}

		first := l.PushBack(1)
		require.Equal(t, first, l.front)
		require.Equal(t, first, l.back)
		require.Nil(t, first.prev)
		require.Nil(t, first.next)

		second := l.PushBack(2)
		require.Equal(t, first, l.front)
		require.Equal(t, second, l.back)
		require.Equal(t, second, first.next)
		require.Equal(t, first, second.prev)
		require.Nil(t, second.next)

		third := l.PushBack(3)
		require.Equal(t, first, l.front)
		require.Equal(t, third, l.back)
		require.Equal(t, third, second.next)
		require.Equal(t, second, third.prev)
		require.Nil(t, third.next)
	})

	t.Run("PopFront", func(t *testing.T) {
		l := &List[int]{}

		l.PushBack(1)
		second := l.PushBack(2)
		third := l.PushBack(3)

		require.Equal(t, 1, l.PopFront())
		require.Equal(t, second, l.front)
		require.Equal(t, third, l.back)

		require.Equal(t, 2, l.PopFront())
		require.Equal(t, third, l.front)
		require.Equal(t, third, l.back)

		require.Equal(t, 3, l.PopFront())
		require.Nil(t, l.front)
		require.Nil(t, l.back)
	})

	t.Run("MoveBack", func(t *testing.T) {
		l := &List[int]{}

		node := l.PushBack(1)
		l.PushBack(2)

		l.MoveBack(node)

		require.Equal(t, 2, l.PopFront())
		require.Equal(t, 1, l.PopFront())
	})

	t.Run("Remove", func(t *testing.T) {
		l := &List[int]{}

		first := l.PushBack(1)
		second := l.PushBack(2)

		l.Remove(first)
		require.Equal(t, second, l.front)
		require.Equal(t, second, l.back)

		l.Remove(second)
		require.Nil(t, l.front)
		require.Nil(t, l.back)
	})

	t.Run("Remove reversed", func(t *testing.T) {
		l := &List[int]{}

		first := l.PushBack(1)
		second := l.PushBack(2)

		l.Remove(second)
		require.Equal(t, first, l.front)
		require.Equal(t, first, l.back)

		l.Remove(first)
		require.Nil(t, l.front)
		require.Nil(t, l.back)
	})
}
