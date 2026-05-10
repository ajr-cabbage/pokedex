package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheData map[string]cacheEntry
	m         sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cData := make(map[string]cacheEntry)
	c := new(Cache)
	c.cacheData = cData
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.m.Lock()
		for key, dat := range c.cacheData {
			t := time.Now()
			if t.Sub(dat.createdAt) >= interval {
				delete(c.cacheData, key)
			}
		}
		c.m.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	defer c.m.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheData[key] = newEntry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	entry, ok := c.cacheData[key]
	if !ok {
		var b []byte
		return b, false
	}
	return entry.val, true
}
