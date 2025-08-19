package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         sync.Mutex
	duration   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(dur time.Duration) *Cache {
	cache := &Cache{
		cacheEntry: make(map[string]cacheEntry),
		duration:   dur,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheEntry[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) >= c.duration {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
