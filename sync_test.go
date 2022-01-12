package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSync(t *testing.T) {
	t.Run("wrapping sync cache returns same instance", func(t *testing.T) {
		s := NewSync(NewBasic[int, int]())
		require.Equal(t, s, NewSync[int, int](s))
	})
}
