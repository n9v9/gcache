package gcache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSync(t *testing.T) {
	t.Run("wrapping sync cache returns same instance", func(t *testing.T) {
		s := NewSync(NewBasic[int, int]())
		require.Equal(t, s, NewSync[int, int](s))
	})

	t.Run("Set", func(t *testing.T) {
		s := NewSync(NewBasic[int, int]())

		limit := 10
		wg := sync.WaitGroup{}
		wg.Add(limit)

		for i := 0; i < limit; i++ {
			go func(i int) {
				defer wg.Done()
				s.Set(i, i)
			}(i)
		}

		wg.Wait()

		for i := 0; i < limit; i++ {
			v, ok := s.Get(i)
			require.True(t, ok)
			require.Equal(t, i, v)
		}
	})

	t.Run("Delete and Len", func(t *testing.T) {
		r := require.New(t)
		s := NewSync(NewBasic[string, int]())

		key := "deletes"
		value := 42
		s.Set(key, value)
		r.Equal(1, s.Len())

		limit := 10
		wg := sync.WaitGroup{}
		wg.Add(limit)

		var deletes int
		var got int
		var mu sync.Mutex

		for i := 0; i < limit; i++ {
			go func() {
				defer wg.Done()
				v, ok := s.Delete(key)

				if ok {
					r.Empty(s.Len())
					mu.Lock()
					deletes++
					got = v
					mu.Unlock()
				}
			}()
		}

		wg.Wait()

		r.Equal(1, deletes)
		r.Equal(value, got)
	})

	t.Run("Do", func(t *testing.T) {
		s := NewSync(NewBasic[string, int]())

		key := "increments"
		limit := 10
		wg := sync.WaitGroup{}
		wg.Add(limit)

		for i := 0; i < limit; i++ {
			go func() {
				defer wg.Done()
				s.Do(func(c Cacher[string, int]) {
					v, _ := c.Get(key)
					c.Set(key, v+1)
				})
			}()
		}

		wg.Wait()

		sum, _ := s.Get(key)
		require.Equal(t, limit, sum)
	})

	t.Run("RDo", func(t *testing.T) {
		s := NewSync(NewBasic[string, int]())

		delay := time.Millisecond * 50
		delta := delay * 2
		limit := 10
		wg := sync.WaitGroup{}
		wg.Add(limit)

		start := time.Now()

		for i := 0; i < limit; i++ {
			go func() {
				defer wg.Done()
				// Check that all RDo calls run concurrently.
				s.RDo(func(_ ReadCacher[string, int]) {
					time.Sleep(delay)
				})
			}()
		}

		wg.Wait()

		taken := time.Since(start)
		require.InDelta(t, delay.Milliseconds(), taken.Milliseconds(), float64(delta.Milliseconds()))
	})
}
