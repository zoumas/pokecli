package cache

import (
	"log"
	"sync"
	"time"
)

type entry struct {
	createdAt time.Time
	data      []byte
}

type Cache struct {
	m  map[string]entry
	mu *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		m:  make(map[string]entry),
		mu: &sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = entry{
		createdAt: time.Now(),
		data:      value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.m[key]
	return v.data, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	t := time.Tick(interval)
	for range t {
		c.mu.Lock()
		for k, v := range c.m {
			if time.Since(v.createdAt) > interval {
				log.Println("Reaped", k)
				delete(c.m, k)
			}
		}
		c.mu.Unlock()
	}
}
