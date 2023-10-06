package cache

import (
	"sync"
	"time"
)

type OnEviction func(map[string]interface{})

type cacheItem struct {
	item       interface{}
	expiration int64
}

type Cache struct {
	mu              sync.RWMutex
	items           map[string]*cacheItem
	ttl             time.Duration
	cleanupInterval time.Duration
	onEvict         OnEviction
	maxItems        int
}

func NewCache(ttl time.Duration, cleanup time.Duration, maxItems int) *Cache {
	return &Cache{
		items: make(map[string]*cacheItem),

		ttl:             ttl,
		cleanupInterval: cleanup,

		maxItems: maxItems,
	}
}

func (c *Cache) Set(key string, item interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.items) == c.maxItems {
		return false
	}

	expiration := time.Now().Add(c.ttl).UnixNano()
	ci := &cacheItem{item, expiration}

	c.items[key] = ci

	return true
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.get(key, false)
}

func (c *Cache) GetAndRefresh(key string) (interface{}, bool) {
	return c.get(key, true)
}

func (c *Cache) get(key string, refresh bool) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]

	if ok {
		if refresh {
			item.expiration = time.Now().Add(c.ttl).UnixNano()
		}

		return item.item, true
	}

	return nil, false
}

func (c *Cache) OnEvict(handler OnEviction) {
	c.onEvict = handler
}

func (c *Cache) Run() {
	go c.run()
}

func (c *Cache) run() {
	ticker := time.NewTicker(c.cleanupInterval)

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		}
	}
}

func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now().UnixNano()

	evicted := make(map[string]interface{})

	for key, item := range c.items {
		if now > item.expiration {
			evicted[key] = item
			delete(c.items, key)
		}
	}

	if c.onEvict != nil && len(evicted) > 0 {
		c.onEvict(evicted)
	}
}
