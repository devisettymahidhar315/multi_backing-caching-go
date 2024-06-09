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

		c.client.LRem(ctx, "cache", 0, key)
	}

	c.client.LPush(ctx, "cache", key)

	c.client.Set(ctx, key, value, 0)

	length, err := c.client.LLen(ctx, "cache").Result()
	if err != nil {
		log.Fatalf("Error getting cache length: %v", err)
	}

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

func (c *LRUCache) Del(key string) {

	exists, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		log.Fatalf("Error checking if key %s exists: %v", key, err)
	}

	if exists == 0 {

		return
	}

	if exists > 0 {

		_, err := c.client.LRem(ctx, "cache", 0, key).Result()
		if err != nil {
			log.Fatalf("Error removing key %s from cache: %v", key, err)
		}

		_, err = c.client.Del(ctx, key).Result()
		if err != nil {
			log.Fatalf("Error deleting key %s: %v", key, err)
		}

	}
}
