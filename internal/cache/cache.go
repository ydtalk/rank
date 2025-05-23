package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

type Cache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

var globalCache *Cache

// InitCache 初始化缓存
func InitCache() {
	globalCache = &Cache{
		items: make(map[string]CacheItem),
	}
}

// Set 设置缓存
func Set(key string, value interface{}, ttl time.Duration) {
	globalCache.mu.Lock()
	defer globalCache.mu.Unlock()

	globalCache.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
}

// Get 获取缓存
func Get(key string) (interface{}, bool) {
	globalCache.mu.RLock()
	defer globalCache.mu.RUnlock()

	item, exists := globalCache.items[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(item.Expiration) {
		delete(globalCache.items, key)
		return nil, false
	}

	return item.Value, true
}

// Delete 删除缓存
func Delete(key string) {
	globalCache.mu.Lock()
	defer globalCache.mu.Unlock()

	delete(globalCache.items, key)
}

// Clear 清空缓存
func Clear() {
	globalCache.mu.Lock()
	defer globalCache.mu.Unlock()

	globalCache.items = make(map[string]CacheItem)
}
