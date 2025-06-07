package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type CacheLocation struct {
	caching  map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func NewLocation(interval time.Duration) *CacheLocation {
	cache := CacheLocation{
		caching:  make(map[string]cacheEntry),
		interval: interval * 30,
	}
	go cache.locationreapLoop()
	return &cache
}

func (cache *CacheLocation) LocationAdd(key string, value []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.caching[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	return
}

func (cache *CacheLocation) LocationGet(key string) (value []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}, found bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.caching[key]
	if ok == true {
		found := true
		return entry.val, found
	}
	return nil, false
}

func (cache *CacheLocation) locationreapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		cache.mu.Lock()
		for key, value := range cache.caching {
			if time.Now().Sub(value.createdAt) > cache.interval {
				fmt.Println("DELETED...")
				delete(cache.caching, key)
			}
		}
		cache.mu.Unlock()
	}
}
