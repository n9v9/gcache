package gocache

// Cacher is the interface that is implemented by every cache.
type Cacher[K comparable, V any] interface {
	// Get returns the value for the given key.
	// If no such key exists then the zero value is returned and ok is false.
	Get(key K) (value V, ok bool)
	// Set inserts a new key value pair into the cache.
	Set(key K, value V)
	// Delete deletes the entry with the given key from the cache and returns
	// the removed value.
	// If no such key exists then the zero value is returned and ok is false.
	Delete(key K) (value V, ok bool)
	// Len returns the number of entries in the cache.
	Len() int
}

type basicCache[K comparable, V any] struct {
	entries map[K]V
}

// NewBasic returns a new cache.
// The cache is not concurrency safe and has no limit for how many entries it
// can keep.
func NewBasic[K comparable, V any]() Cacher[K, V] {
	return &basicCache[K, V]{
		entries: make(map[K]V),
	}
}

func (b *basicCache[K, V]) Get(key K) (value V, ok bool) {
	value, ok = b.entries[key]
	return
}

func (b *basicCache[K, V]) Set(key K, value V) {
	b.entries[key] = value
}

func (b *basicCache[K, V]) Delete(key K) (value V, ok bool) {
	value, ok = b.entries[key]
	delete(b.entries, key)
	return
}

func (b *basicCache[K, V]) Len() int {
	return len(b.entries)
}
