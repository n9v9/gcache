package gocache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		c := NewBasic[int, string]()

		v, ok := c.Get(1)
		require.Equal(t, "", v)
		require.Equal(t, false, ok)

		c.Set(1, "A")

		v, ok = c.Get(1)
		require.Equal(t, "A", v)
		require.Equal(t, true, ok)
	})

	t.Run("delete", func(t *testing.T) {
		c := NewBasic[int, string]()

		v, ok := c.Delete(1)
		require.Equal(t, "", v)
		require.Equal(t, false, ok)

		c.Set(1, "A")

		v, ok = c.Delete(1)
		require.Equal(t, "A", v)
		require.Equal(t, true, ok)
	})

	t.Run("len", func(t *testing.T) {
		c := NewBasic[int, string]()
		require.Equal(t, 0, c.Len())

		c.Set(1, "A")
		require.Equal(t, 1, c.Len())

		c.Set(2, "B")
		require.Equal(t, 2, c.Len())

		c.Delete(2)
		require.Equal(t, 1, c.Len())

		c.Delete(1)
		require.Equal(t, 0, c.Len())
	})
}
