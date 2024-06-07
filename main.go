package main

import (
	"app/multi_cache"
	"time"
)

const length = 3

func main() {

	cache := multi_cache.NewMultiCache()

	cache.Set("key1", "value1", length)
	cache.Set("key2", "value2", length)
	cache.Set("key3", "value3", length)
	time.Sleep(1 * time.Second)
	cache.Print()

	go cache.Get("key2")
	time.Sleep(1 * time.Second)
	cache.Print()

	go cache.Set("key3", "value4", length)
	time.Sleep(1 * time.Second)
	cache.Print()

	go cache.Set("key4", "value5", length)
	time.Sleep(1 * time.Second)
	cache.Print()
}
