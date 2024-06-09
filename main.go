package main

import (
	"app/multi_cache"
	"fmt"
	"sync"
)

const length = 3

func main() {
	cache := multi_cache.NewMultiCache()
	var wg sync.WaitGroup

	fmt.Println("Setting key1, key2, key3")
	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Set("key1", "value1", length)
		cache.Set("key2", "value2", length)
		cache.Set("key3", "value3", length)
	}()
	wg.Wait()
	cache.Print()

	fmt.Println("Getting key2")
	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Get("key2")
	}()
	wg.Wait()
	cache.Print()

	fmt.Println("Setting key3 to new value4")
	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Set("key3", "value4", length)
	}()
	wg.Wait()
	cache.Print()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Del("key3")
	}()
	wg.Wait()
	cache.Print()
}
