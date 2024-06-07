package multi_cache

import (
	"app/in_memory"
	"app/redis"
	"sync"
)

type MultiCache struct {
	redisCache    *redis.LRUCache
	inMemoryCache *in_memory.LRUCache
}

func NewMultiCache() *MultiCache {
	return &MultiCache{
		redisCache:    redis.NewLRUCache(),
		inMemoryCache: in_memory.NewLRUCache(),
	}
}

func (c *MultiCache) Set(key, value string, length int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.redisCache.Put(key, value, length)
	}()
	go func() {
		defer wg.Done()
		c.inMemoryCache.Put(key, value, length)
	}()
	wg.Wait()
}

func (c *MultiCache) Get(key string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.inMemoryCache.Get(key)
	}()
	go func() {
		defer wg.Done()
		c.redisCache.Get(key)
	}()
	wg.Wait()
}

func (c *MultiCache) Print() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.redisCache.Print()
	}()
	go func() {
		defer wg.Done()
		c.inMemoryCache.Print()
	}()
	wg.Wait()
}
