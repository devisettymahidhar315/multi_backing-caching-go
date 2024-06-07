package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type LRUCache struct {
	client *redis.Client
}

func NewLRUCache() *LRUCache {
	opts := &redis.Options{
		Addr:     "localhost:6379",
		PoolSize: 10,
	}

	rdb := redis.NewClient(opts)

	rdb.Del(ctx, "cache")

	return &LRUCache{
		client: rdb,
	}
}

func (c *LRUCache) Put(key, value string, maxLength int) {

	exists, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		log.Fatalf("Error checking if key %s exists: %v", key, err)
	}
	if exists > 0 {
		// Move the key to the front of the list
		c.client.LRem(ctx, "cache", 0, key)
	}

	// Add the key to the front of the list
	c.client.LPush(ctx, "cache", key)
	// Set the value for the key
	c.client.Set(ctx, key, value, 0)

	// Get the current length of the cache list
	length, err := c.client.LLen(ctx, "cache").Result()
	if err != nil {
		log.Fatalf("Error getting cache length: %v", err)
	}
	// If the cache length exceeds the max length, remove the oldest key
	if length > int64(maxLength) {
		oldest, err := c.client.RPop(ctx, "cache").Result()
		if err != nil {
			log.Fatalf("Error popping oldest key: %v", err)
		}
		c.client.Del(ctx, oldest)
	}
}

func (c *LRUCache) Get(key string) (string, bool) {

	value, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		log.Fatalf("Error getting key %s: %v", key, err)
	}

	c.client.LRem(ctx, "cache", 0, key)
	c.client.LPush(ctx, "cache", key)

	return value, true
}

func (c *LRUCache) Print() {

	fmt.Println("redis")
	keys, err := c.client.LRange(ctx, "cache", 0, -1).Result()
	if err != nil {
		log.Fatalf("Error getting cache keys: %v", err)
	}

	for _, key := range keys {

		value, err := c.client.Get(ctx, key).Result()
		if err != nil {
			log.Fatalf("Error getting key %s: %v", key, err)
		}
		fmt.Printf("%s: %s\n", key, value)
	}
}
