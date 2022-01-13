package gcache

// ReacCacher is the interface that is implemented by every read only cache.
type ReadCacher[K comparable, V any] interface {
	// Get returns the value for the given key.
	// If no such key exists then the zero value is returned and ok is false.
	Get(key K) (value V, ok bool)
	// Len returns the number of entries in the cache.
	Len() int
	// Keys returns a slice of all keys in the cache.
	Keys() []K
	// Values returns a slice of all values in the cache.
	Values() []V
	// ForEach calls the given function for each entry in the cache.
	ForEach(func(key K, value V))
}

// Cacher is the interface that is implemented by every read and write cache.
type Cacher[K comparable, V any] interface {
	ReadCacher[K, V]
	// Set inserts a new key value pair into the cache.
	Set(key K, value V)
	// Delete deletes the entry with the given key from the cache and returns
	// the removed value.
	// If no such key exists then the zero value is returned and ok is false.
	Delete(key K) (value V, ok bool)
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

func (b *basicCache[K, V]) Keys() []K {
	keys := make([]K, 0, len(b.entries))
	for key := range b.entries {
		keys = append(keys, key)
	}
	return keys
}

func (b *basicCache[K, V]) Values() []V {
	values := make([]V, 0, len(b.entries))
	for _, value := range b.entries {
		values = append(values, value)
	}
	return values
}

func (b *basicCache[K, V]) ForEach(f func(K, V)) {
	for key, value := range b.entries {
		f(key, value)
	}
}
