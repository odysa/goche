package cache

type Key = []byte
type Value = []byte

type Cache interface {
	get(key Key) (Value, error)
	set(key Key, value Value) (bool, error)
}
