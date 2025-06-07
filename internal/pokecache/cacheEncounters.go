package pokecache

import (
	"sync"
	"time"
)

type CacheEncounters struct {
	caching  map[string]cachePokemons
	interval time.Duration
	mu       sync.Mutex
}

type cachePokemons struct {
	createdAt time.Time
	val       []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	}
}

func NewEncounter(interval time.Duration) *CacheEncounters {
	cache := CacheEncounters{
		caching:  make(map[string]cachePokemons),
		interval: interval * 30,
	}
	go cache.encounterReapLoop()
	return &cache
}

func (cache *CacheEncounters) EncounterAdd(key string, value []struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int   `json:"chance"`
			ConditionValues []any `json:"condition_values"`
			MaxLevel        int   `json:"max_level"`
			Method          struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"method"`
			MinLevel int `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int `json:"max_chance"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
}) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.caching[key] = cachePokemons{
		createdAt: time.Now(),
		val:       value,
	}
	return
}

func (cache *CacheEncounters) EncounterGet(key string) (value []struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int   `json:"chance"`
			ConditionValues []any `json:"condition_values"`
			MaxLevel        int   `json:"max_level"`
			Method          struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"method"`
			MinLevel int `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int `json:"max_chance"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
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

func (cache *CacheEncounters) encounterReapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		cache.mu.Lock()
		for key, value := range cache.caching {
			if time.Now().Sub(value.createdAt) > cache.interval {
				delete(cache.caching, key)
			}
		}
		cache.mu.Unlock()
	}
}
