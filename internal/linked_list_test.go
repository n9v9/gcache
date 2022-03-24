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
		r := require.New(t)
		l := &List[int]{}

		first := l.PushBack(1)
		r.Equal(first, l.front)
		r.Equal(first, l.back)
		r.Nil(first.prev)
		r.Nil(first.next)

		second := l.PushBack(2)
		r.Equal(first, l.front)
		r.Equal(second, l.back)
		r.Equal(second, first.next)
		r.Equal(first, second.prev)
		r.Nil(second.next)

		third := l.PushBack(3)
		r.Equal(first, l.front)
		r.Equal(third, l.back)
		r.Equal(third, second.next)
		r.Equal(second, third.prev)
		r.Nil(third.next)
	})

	t.Run("PopFront", func(t *testing.T) {
		r := require.New(t)
		l := &List[int]{}

		l.PushBack(1)
		second := l.PushBack(2)
		third := l.PushBack(3)

		r.Equal(1, l.PopFront())
		r.Equal(second, l.front)
		r.Equal(third, l.back)

		r.Equal(2, l.PopFront())
		r.Equal(third, l.front)
		r.Equal(third, l.back)

		r.Equal(3, l.PopFront())
		r.Nil(l.front)
		r.Nil(l.back)
	})

	t.Run("MoveBack", func(t *testing.T) {
		r := require.New(t)
		l := &List[int]{}

		node := l.PushBack(1)
		l.PushBack(2)

		l.MoveBack(node)

		r.Equal(2, l.PopFront())
		r.Equal(1, l.PopFront())
	})

	t.Run("Remove", func(t *testing.T) {
		r := require.New(t)
		l := &List[int]{}

		first := l.PushBack(1)
		second := l.PushBack(2)

		l.Remove(first)
		r.Equal(second, l.front)
		r.Equal(second, l.back)

		l.Remove(second)
		r.Nil(l.front)
		r.Nil(l.back)
	})

	t.Run("Remove reversed", func(t *testing.T) {
		r := require.New(t)
		l := &List[int]{}

		first := l.PushBack(1)
		second := l.PushBack(2)

		l.Remove(second)
		r.Equal(first, l.front)
		r.Equal(first, l.back)

		l.Remove(first)
		r.Nil(l.front)
		r.Nil(l.back)
	})
}
