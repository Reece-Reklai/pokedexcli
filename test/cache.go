package test

import (
	"sync"
	"time"
)

type Cache struct {
	caching  map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		caching:  make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.ReapLoop()
	return &cache
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.caching[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	return
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	value, ok := cache.caching[key]
	if ok == true {
		return value.val, true
	} else {
		return nil, false
	}
}

func (cache *Cache) ReapLoop() {
	for {
		ticker := time.NewTicker(cache.interval)
		defer ticker.Stop()
		reciever := <-ticker.C
		for key, value := range cache.caching {
			if reciever.Compare(value.createdAt) == 1 {
				delete(cache.caching, key)
			}
		}
	}
}
