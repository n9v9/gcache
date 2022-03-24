# gcache

[![Go Reference](https://pkg.go.dev/badge/github.com/n9v9/gcache.svg)](https://pkg.go.dev/github.com/n9v9/gcache)
[![Go Report Card](https://goreportcard.com/badge/github.com/n9v9/gcache)](https://goreportcard.com/report/github.com/n9v9/gcache)

gcache is a generic caching library for Go 1.18+.

## Install

```
go get github.com/n9v9/gcache
```

## Concepts

Caches can be mixed and matched to build different types of caches.

All cache types are represented by these three interfaces:

-   `ReadCacher`: Read-only access to a cache
-   `Cacher`: Read and write access to a cache
-   `SyncCacher`: Concurrency safe implementation around a `Cacher`

## Examples

### Basic Cache

```go
cache := gcache.NewBasic[string, int]()

cache.Set("visitors", 1)

// Get the number of entries in the cache.
l := cache.Len()

// 1, true because the key exists.
visitors, ok := cache.Get("visitors")

// 1, true because the key exists.
visitors, ok = cache.Delete("visitors")

// zero value, false because the key does not exist anymore.
visitors, ok = cache.Get("visitors")

// zero value, false because the key does not exist anymore.
visitors, ok = cache.Delete("visitors")

// Get keys and values from the cache.
keys := cache.Keys()
values := cache.Values()

// Iterate over each key value pair in the cache.
cache.ForEach(func(key string, value int) {
    // ...
})
```

### LRU Cache

```go
// Limit the cache to a maximum size of 3 entries.
cache := gcache.NewLRU[string, int](3)

cache.Set("A", 1)
cache.Set("B", 2)
cache.Set("C", 3)

// After this call to Set, "A" will be removed.
cache.Set("D", 4)

// Access "B" so it becomes the most recently used entry.
cache.Get("B")

// After this call to Set, "C" will be removed because the max size of the cache
// is reached and we accessed "B" most recently and thus "C" becomes the least
// recently used entry.
cache.Set("E", 5)
```

### Sync Cache

```go
// Create a concurrency safe LRU cache.
cache := gcache.NewSync(gcache.NewLRU[string, int](3))

cache.Do(func(c gcache.Cacher[string, int]) {
    // Exclusive read-write access to the cache via c.
    visitors, _ := c.Get("visitors")
    c.Set("visitors", visitors+1)
})

cache.RDo(func(c gcache.ReadCacher[string, int]) {
    // Concurrent read-only access to the cache via c.
})
```
