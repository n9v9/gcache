package gcache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		r := require.New(t)
		c := NewBasic[int, string]()

		v, ok := c.Get(1)
		r.Empty(v)
		r.False(ok)

		c.Set(1, "A")

		v, ok = c.Get(1)
		r.Equal("A", v)
		r.True(ok)
	})

	t.Run("delete", func(t *testing.T) {
		r := require.New(t)
		c := NewBasic[int, string]()

		v, ok := c.Delete(1)
		r.Empty(v)
		r.False(ok)

		c.Set(1, "A")

		v, ok = c.Delete(1)
		r.Equal("A", v)
		r.True(ok)
	})

	t.Run("len", func(t *testing.T) {
		r := require.New(t)
		c := NewBasic[int, string]()
		r.Empty(c.Len())

		c.Set(1, "A")
		r.Equal(1, c.Len())

		c.Set(2, "B")
		r.Equal(2, c.Len())

		c.Delete(2)
		r.Equal(1, c.Len())

		c.Delete(1)
		r.Equal(0, c.Len())
	})

	t.Run("keys and values", func(t *testing.T) {
		r := require.New(t)
		c := NewBasic[int, string]()

		c.Set(1, "A")
		c.Set(2, "B")
		c.Set(3, "C")

		r.ElementsMatch([]int{1, 2, 3}, c.Keys())
		r.ElementsMatch([]string{"A", "B", "C"}, c.Values())
	})
}
