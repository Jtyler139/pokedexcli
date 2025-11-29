package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	data		map[string]cacheEntry
	mu			sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	dat := make(map[string]cacheEntry)
	
	c := Cache{
		data: 	dat,
		mu: 	sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: 	time.Now(),
		val: 		value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for t := range ticker.C {
		c.Reap(t, interval)
	}
}

func (c *Cache) Reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	for key, entry := range c.data {
		if now.Sub(entry.createdAt) > interval {
			delete(c.data, key)
		}
	}
	c.mu.Unlock()
}