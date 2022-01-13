package gocache

import "github.com/n9v9/gocache/internal"

type lruEntry[K comparable, V any] struct {
	node  *internal.Node[K]
	value V
}

type lruCache[K comparable, V any] struct {
	cache   Cacher[K, *lruEntry[K, V]]
	keys    *internal.List[K]
	maxSize int
}

// NewLRU returns a new cache that removes the least recently used entry when
// maxSize is reached.
// The function panics when maxSize is less than 1.
func NewLRU[K comparable, V any](maxSize int) Cacher[K, V] {
	if maxSize <= 0 {
		panic("cache: maxSize must be greater than zero")
	}
	return &lruCache[K, V]{
		cache:   NewBasic[K, *lruEntry[K, V]](),
		keys:    &internal.List[K]{},
		maxSize: maxSize,
	}
}

func (l *lruCache[K, V]) Get(key K) (value V, ok bool) {
	entry, ok := l.cache.Get(key)

	if !ok {
		var zero V
		return zero, false
	}

	l.keys.MoveBack(entry.node)

	return entry.value, true
}

func (l *lruCache[K, V]) Set(key K, value V) {
	entry, ok := l.cache.Get(key)

	if ok {
		// This modifies the value already in the cache, so there
		// is no need for another Set on the inner cache.
		entry.value = value
		l.keys.MoveBack(entry.node)
	} else {
		node := l.keys.PushBack(key)
		entry = &lruEntry[K, V]{
			node:  node,
			value: value,
		}
		l.cache.Set(key, entry)
		if l.cache.Len() > l.maxSize {
			old := l.keys.PopFront()
			l.cache.Delete(old)
		}
	}
}

func (l *lruCache[K, V]) Delete(key K) (value V, ok bool) {
	entry, ok := l.cache.Delete(key)

	if !ok {
		var zero V
		return zero, false
	}

	l.keys.Remove(entry.node)

	return entry.value, true
}

func (l *lruCache[K, V]) Len() int {
	return l.cache.Len()
}

func (l *lruCache[K, V]) Keys() []K {
	return l.cache.Keys()
}

func (l *lruCache[K, V]) Values() []V {
	values := make([]V, 0, l.cache.Len())
	l.cache.ForEach(func(_ K, value *lruEntry[K, V]) {
		values = append(values, value.value)
	})
	return values
}

func (l *lruCache[K, V]) ForEach(f func(K, V)) {
	l.cache.ForEach(func(key K, value *lruEntry[K, V]) {
		f(key, value.value)
	})
}
