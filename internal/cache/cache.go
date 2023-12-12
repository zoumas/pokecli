package cache

import (
	"sync"
	"time"
)

type entry struct {
	createdAt time.Time
	payload   []byte
}

type Cache struct {
	m            map[string]entry
	mu           *sync.RWMutex
	reapInterval time.Duration
}

func NewCache(reapInterval time.Duration) *Cache {
	c := &Cache{
		m:            make(map[string]entry),
		reapInterval: reapInterval,
		mu:           &sync.RWMutex{},
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = entry{
		createdAt: time.Now(),
		payload:   value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.m[key]
	return v.payload, ok
}

func (c *Cache) reapLoop() {
	for range time.Tick(c.reapInterval) {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.m {
		if time.Since(v.createdAt) > c.reapInterval {
			delete(c.m, k)
		}
	}
}
