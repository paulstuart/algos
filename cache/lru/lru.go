package lru

import (
	"container/list"
	"fmt"
	"sync"
)

type kvpair struct {
	Key, Value string
}

// LRU is a Least Recently Used cache
type LRU struct {
	mu      sync.Mutex
	entries map[string]*list.Element
	max     int
	last    *list.List
}

// NewLRU returns a LRU
func NewLRU(size int) *LRU {
	return &LRU{
		max:     size,
		entries: make(map[string]*list.Element, size),
		last:    list.New(),
	}
}

// Get returns a key value
func (u *LRU) Get(key string) (string, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if v, ok := u.entries[key]; ok {
		u.last.MoveToBack(v)
		return v.Value.(kvpair).Value, true
	}
	return "", false
}

// Getter will return the cached value or will
// update the cache with the fn results and return that
func (u *LRU) Getter(key string, fn func(string) string) string {
	// TODO: singlefile map of mutexes for more granular locking?
	u.mu.Lock()

	if v, ok := u.entries[key]; ok {
		u.last.MoveToBack(v)
		u.mu.Unlock()
		return v.Value.(kvpair).Value
	}
	u.mu.Unlock()
	value := fn(key)
	u.Set(key, value)
	return value
}

// Set saves the key/value in the cache
func (u *LRU) Set(key string, value string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	// no eviction required?
	if len(u.entries) < u.max {
		u.entries[key] = u.last.PushBack(kvpair{key, value})
		return
	}

	// last touched should be at the front of the list
	old := u.last.Front()
	delete(u.entries, old.Value.(kvpair).Key)
	u.last.Remove(old)
	u.entries[key] = u.last.PushBack(kvpair{key, value})

}

// show is a debug dump of cache contents
func (u *LRU) show() {
	fmt.Println("MAP:")
	for k, v := range u.entries {
		vv := v.Value.(kvpair)
		fmt.Printf("K:%20s K2: %20s V:%20s\n", k, vv.Key, vv.Value)
	}
	fmt.Println("LIST:")
	for elem := u.last.Front(); elem != nil; elem = elem.Next() {
		vv := elem.Value.(kvpair)
		fmt.Printf("kvpair:%+v\n", vv)
	}
}
