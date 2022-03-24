package gcache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	t.Run("non positive max size panics", func(t *testing.T) {
		require.Panics(t, func() { NewLRU[int, int](0) })
		require.Panics(t, func() { NewLRU[int, int](-1) })
	})

	entryExists := func(t *testing.T, c Cacher[string, int], key string, value int) {
		v, ok := c.Get(key)
		require.True(t, ok)
		require.Equal(t, value, v)
	}

	entryNotExists := func(t *testing.T, c Cacher[string, int], key string) {
		_, ok := c.Get(key)
		require.False(t, ok)
	}

	t.Run("different keys", func(t *testing.T) {
		l := NewLRU[string, int](2)

		l.Set("A", 1)
		entryExists(t, l, "A", 1)

		l.Set("B", 2)
		entryExists(t, l, "A", 1)
		entryExists(t, l, "B", 2)

		l.Set("C", 3)
		entryNotExists(t, l, "A")
	})

	t.Run("renew by overwriting entry", func(t *testing.T) {
		l := NewLRU[string, int](2)

		l.Set("A", 1)
		l.Set("B", 2)
		l.Set("A", 10)
		l.Set("C", 3)

		entryNotExists(t, l, "B")
		entryExists(t, l, "A", 10)
		entryExists(t, l, "C", 3)
	})

	t.Run("renew by getting entry", func(t *testing.T) {
		l := NewLRU[string, int](2)

		l.Set("A", 1)
		l.Set("B", 2)

		l.Get("A")
		l.Set("C", 3)

		entryNotExists(t, l, "B")
		entryExists(t, l, "A", 1)
		entryExists(t, l, "C", 3)
	})

	t.Run("delete and len", func(t *testing.T) {
		r := require.New(t)
		l := NewLRU[string, int](2)

		r.Empty(l.Len())

		_, ok := l.Delete("A")
		r.False(ok)

		l.Set("A", 1)
		l.Set("B", 2)
		r.Equal(2, l.Len())

		v, ok := l.Delete("A")
		r.True(ok)
		r.Equal(1, v)
		r.Equal(1, l.Len())
		entryNotExists(t, l, "A")
		entryExists(t, l, "B", 2)

		v, ok = l.Delete("B")
		r.True(ok)
		r.Equal(2, v)
		r.Equal(0, l.Len())
		entryNotExists(t, l, "B")
	})
}
