package cache

type Cacher[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Delete(key K) (V, bool)
	Len() int
}

type basicCache[K comparable, V any] struct {
	entries map[K]V
}

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
