package pokecache

import (
	"fmt"
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
	// kstr := "key" + time.Now().Format(time.UnixDate)
	// fmt.Println("Adding...", kstr)
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		fmt.Println("couldn't find key: ", key)
		return []byte{}, false
	}

	// fmt.Println("Found key: ", key, entry.val, entry.createdAt)
	//entry and it's found
	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	// ticker := time.NewTicker(interval * time.Second)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer c.mu.Unlock()
	// c.Add("new", []byte("howdy"))
	// fmt.Println("reaping... maybe")
	for {
		select {
		case timeNow := <-ticker.C:
			// fmt.Println("timeNow", timeNow)
			// c.Get("new")
			c.mu.Lock()
			for key, entry := range c.data {
				fmt.Println("entry + interval", entry.createdAt.Add(interval))
				isOld := entry.createdAt.Add(interval).Before(timeNow)
				if isOld {
					delete(c.data, key)
					fmt.Println("deleted key: ", key)
				}
			}
			c.mu.Unlock()
			// fmt.Println("end of case")
		}
		// fmt.Println("end of select")
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: map[string]cacheEntry{
			// 	"key1": {
			// 		createdAt: time.Now().Add(-3 * time.Minute),
			// 		val:       []byte("value1"),
			// 	},
			// 	"key2": {
			// 		createdAt: time.Now().Add(5 * time.Second),
			// 		val:       []byte("value1"),
			// 	},
		},
		mu: new(sync.Mutex), // Proper mutex initialization
	}
	go cache.reapLoop(interval)
	return cache
}
