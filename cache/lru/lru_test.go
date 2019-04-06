package lru

import (
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU(3)
	lru.Set("1", "one")
	lru.Set("2", "two")
	lru.Set("3", "three")
	lru.Set("4", "foure")
	lru.Set("5", "five")
	lru.show()
}

func TestExpire(t *testing.T) {
	lru := NewLRU(5)
	lru.Set("1", "one")
	lru.Set("2", "two")
	lru.Set("3", "three")
	lru.Set("4", "foure")
	lru.Set("5", "five")
	lru.Get("1")
	lru.Get("2")
	lru.Get("3")
	lru.Get("4")
	lru.Get("1")
	lru.Get("2")
	lru.Get("3")
	lru.Get("4")
	lru.Set("99", "not five")

	lru.show()
}
