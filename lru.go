package cache

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
		l.keys.MoveBack(entry.node)
	} else {
		node := l.keys.PushBack(key)
		entry = &lruEntry[K, V]{
			node:  node,
			value: value,
		}
		if l.cache.Len() > l.maxSize {
			l.keys.PopFront()
		}
	}

	l.cache.Set(key, entry)
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
