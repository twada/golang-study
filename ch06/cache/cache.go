package cache

import "sync"

var cache = struct {
	sync.Mutex
	mapping map[string]string
} {
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func Set(key string, val string) {
	cache.Lock()
	cache.mapping[key] = val
	cache.Unlock()
}
