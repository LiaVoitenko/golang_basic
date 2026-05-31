package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]item
	mu   sync.RWMutex
}

type item struct {
	value        any
	expirationAt time.Time
}

func New() *Cache {
	return &Cache{
		data: make(map[string]item),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	expiration := now.Add(ttl)
	c.data[key] = item{
		value:        value,
		expirationAt: expiration,
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	value, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	if time.Now().After(value.expirationAt) {
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return nil, false
	}
	return value.value, ok
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}
