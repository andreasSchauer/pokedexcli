package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries		map[string]cacheEntry
	mu			*sync.Mutex
}


type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}


func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return cache
}


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: 	time.Now(),
		val:		val,
	}
}


func (c *Cache) Get(key string) (val []byte, keyExists bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if entry, exists := c.entries[key]; exists {
		return entry.val, true
	}

	return nil, false
}


func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}


func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if entry.isExpired(interval) {
			delete(c.entries, key)
		}
	}
}


func (ce *cacheEntry) isExpired (interval time.Duration) bool {
	return ce.createdAt.Add(interval).Before(time.Now())
}