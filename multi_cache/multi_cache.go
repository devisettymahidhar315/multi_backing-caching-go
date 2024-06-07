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

	var wg1 sync.WaitGroup
	wg1.Add(2)
	go c.redisCache.Put(key, value, length, &wg1)
	go c.inMemoryCache.Put(key, value, length, &wg1)
	wg1.Wait()
}

func (c *MultiCache) Get(key string) {

	var wg1 sync.WaitGroup
	wg1.Add(2)
	go c.inMemoryCache.Get(key, &wg1)
	go c.redisCache.Get(key, &wg1)
	wg1.Wait()

}

func (c *MultiCache) Print() {

	var wg1 sync.WaitGroup
	wg1.Add(2)
	go c.redisCache.Print(&wg1)
	go c.inMemoryCache.Print(&wg1)
	wg1.Wait()
}
