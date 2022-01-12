package cache

import "sync"

type Syncer[K comparable, V any] interface {
	Cacher[K, V]
	Do(func())
}

type syncCache[K comparable, V any] struct {
	mu    *sync.RWMutex
	cache Cacher[K, V]
}

func NewSync[K comparable, V any](c Cacher[K, V]) Syncer[K, V] {
	if c, ok := c.(*syncCache[K, V]); ok {
		return c
	}

	return &syncCache[K, V]{
		mu:    &sync.RWMutex{},
		cache: c,
	}
}

func (s *syncCache[K, V]) Get(key K) (value V, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cache.Get(key)
}

func (s *syncCache[K, V]) Set(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache.Set(key, value)
}

func (s *syncCache[K, V]) Delete(key K) (value V, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.cache.Delete(key)
}

func (s *syncCache[K, V]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cache.Len()
}

func (s *syncCache[K, V]) Do(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f()
}
