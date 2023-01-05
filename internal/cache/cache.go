package cache

type Size int

type Value interface {
}

type Cache[K comparable, V any] interface {
	Get(key K) (*V, error)
	Set(key K, value V) error
}
