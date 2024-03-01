package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer c.mu.Unlock()
	for {
		select {
		case timeNow := <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.data {
				isOld := entry.createdAt.Add(interval).Before(timeNow)
				if isOld {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

// Or:
// func (c *Cache) reapLoopAlt(interval time.Duration) {
// 	ticker := time.NewTicker(interval)
// 	for range ticker.C { // instead of the for select idiom
// 		c.reap(time.Now(), interval)
// 	}
// }
//
// func (c *Cache) reap(now time.Time, last time.Duration) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	for k, v := range c.data {
// 		if v.createdAt.Before(now.Add(-last)) {
// 			delete(c.data, k)
// 		}
// 	}
// }

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: make(map[string]cacheEntry),
		// mu:   new(sync.Mutex), // Proper mutex initialization
		mu: &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}
