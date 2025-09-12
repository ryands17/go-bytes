package utils

import (
	"errors"
	"sync"
	"time"
)

type (
	item struct {
		value any
		ttl   int64
	}

	cache struct {
		cacheMap   map[string]item
		quit       chan struct{}
		mx         sync.RWMutex
		defaultTTL time.Duration
	}
)

func NewCache(defaultTTL time.Duration, cleanupInterval *time.Duration) *cache {
	if cleanupInterval == nil {
		cleanupInterval = PointerTo(3 * time.Minute)
	}

	c := &cache{
		cacheMap:   make(map[string]item),
		quit:       make(chan struct{}),
		defaultTTL: defaultTTL,
	}

	// background cleaner
	go func() {
		ticker := time.NewTicker(*cleanupInterval)
		for {
			select {
			case <-ticker.C:
				c.cleanup()
			case <-c.quit:
				ticker.Stop()
				return
			}
		}
	}()

	return c
}

func (c *cache) Set(key string, value any, ttl *time.Duration) error {
	if ttl == nil {
		ttl = &c.defaultTTL
	}

	if key == "" || value == nil {
		return errors.New("key and value must be non-empty")
	}

	c.mx.Lock()
	defer c.mx.Unlock()

	c.cacheMap[key] = item{
		value: value,
		ttl:   time.Now().Add(*ttl).UnixNano(),
	}

	return nil
}

func (c *cache) Get(key string) (any, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	it, exists := c.cacheMap[key]
	if !exists || (it.ttl > 0 && time.Now().UnixNano() > it.ttl) {
		return nil, false
	}

	return it.value, true
}

func (c *cache) Delete(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()

	delete(c.cacheMap, key)
}

func (c *cache) cleanup() {
	c.mx.Lock()
	defer c.mx.Unlock()

	now := time.Now().UnixNano()
	for k, v := range c.cacheMap {
		if v.ttl > 0 && now > v.ttl {
			delete(c.cacheMap, k)
		}
	}
}
