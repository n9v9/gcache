package gocache

import "sync"

// Syncer is the interface that is implemented by concurrency safe cache
// implementaitons.
type Syncer[K comparable, V any] interface {
	Cacher[K, V]
	// Do executes the given function in an atomic context giving it exclusive
	// access to the cache.
	// That means that while Do is executing, read and write access to the cache
	// is locked.
	Do(func(Cacher[K, V]))
	// RDo executes the given function in an atomic context giving it exclusive
	// read access to the cache.
	// That means that while RDo is executing, write access to the cache is
	// locked.
	RDo(func(ReadCacher[K, V]))
}

type syncCache[K comparable, V any] struct {
	mu    *sync.RWMutex
	cache Cacher[K, V]
}

// NewSync makes c safe for concurrent access.
// If c is already a Syncer, then this function is a no-op.
func NewSync[K comparable, V any](c Cacher[K, V]) Syncer[K, V] {
	if c, ok := c.(Syncer[K, V]); ok {
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

func (s *syncCache[K, V]) Keys() []K {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cache.Keys()
}

func (s *syncCache[K, V]) Values() []V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cache.Values()
}

func (s *syncCache[K, V]) ForEach(f func(K, V)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache.ForEach(f)
}

func (s *syncCache[K, V]) Do(f func(Cacher[K, V])) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f(s.cache)
}

func (s *syncCache[K, V]) RDo(f func(ReadCacher[K, V])) {
	s.mu.RLock()
	defer s.mu.RLock()
	f(s.cache)
}
