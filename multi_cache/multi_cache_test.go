package multi_cache

import (
	"fmt"
	"testing"
)

// Define the capacity of the cache
const len = 2

// unit testing
// TestGET tests the Get method of the cache
func TestGET(t *testing.T) {
	// Create a new multi-cache instance
	cache := NewMultiCache()
	// Set key-value pairs in the cache
	cache.Set("a", "1", len)
	cache.Set("b", "2", len)

	// Test case 1: Get a non-existent key
	res1 := cache.Get("c")
	if res1 != "" {
		t.Error("case 1 error: expected empty string for non-existent key")
	}

	// Test case 2: Get an existing key
	res2 := cache.Get("a")
	if res2 != "1" {
		t.Error("case 2 error: expected '1' for key 'a'")
	}
}

// TestPrint tests the Print methods of the cache
func TestPrint(t *testing.T) {
	// Create a new multi-cache instance
	cache := NewMultiCache()
	// Set key-value pairs in the cache
	cache.Set("a", "1", len)
	cache.Set("b", "2", len)

	// Get the printed results from both in-memory and Redis caches
	inmemory_result := cache.Print_in_mem()
	redis_result := cache.Print_redis()
	// Check if the data is the same in both backends
	if inmemory_result != redis_result {
		t.Error("data is not the same in both backends")
	}

	// Set another key-value pair to exceed the cache capacity
	cache.Set("c", "3", len)
	inmemory_result = cache.Print_in_mem()
	redis_result = cache.Print_redis()
	expected_output := "c:3, b:2"

	// Check if the expected output matches the in-memory result
	if inmemory_result != expected_output {
		t.Error("in-memory cache does not match expected output")

	}
	// Check if the expected output matches the Redis result
	fmt.Println(redis_result, expected_output)
	if redis_result != expected_output {
		t.Error("Redis cache does not match expected output")
	}
}

// TestDel tests the Del method of the cache
func TestDel(t *testing.T) {
	// Create a new multi-cache instance
	cache := NewMultiCache()
	// Set key-value pairs in the cache
	cache.Set("a", "1", len)
	cache.Set("b", "2", len)

	// Delete a key from the cache
	cache.Del("a")
	// Check if the deleted key returns an empty string
	result := cache.Get("a")
	if result != "" {
		t.Error("expected empty string for deleted key 'a'")
	}
	// Check if an existing key returns the correct value
	result = cache.Get("b")
	if result != "2" {
		t.Error("expected '2' for key 'b'")
	}
}

//benchmarking
//becnchmarking for set method of the cache

func BenchmarkSet(b *testing.B) {
	// Create a new multi-cache instance
	cache := NewMultiCache()
	for i := 0; i < b.N; i++ {
		cache.Set("a", "1", len)
	}
}

//benchmarking for get methods of the cache

func BenchmarkGet(b *testing.B) {
	// Create a new multi-cache instance
	cache := NewMultiCache()

	for i := 0; i < b.N; i++ {
		cache.Get("a")

	}
}

//benchmarking for del methods of the cache

func BenchmarkDel(b *testing.B) {
	// Create a new multi-cache instance

	cache := NewMultiCache()
	cache.Set("a", "1", len)
	for i := 0; i < b.N; i++ {
		cache.Del("a")

	}
}

//benchmarking for printing redis methods of the cache

func BenchmarkPrintRedis(b *testing.B) {
	// Create a new multi-cache instance

	cache := NewMultiCache()
	cache.Set("a", "1", len)
	for i := 0; i < b.N; i++ {
		cache.Print_redis()

	}
}

//benchmarking for printing in memory methods of the cache

func BenchmarkPrintinmemory(b *testing.B) {
	// Create a new multi-cache instance

	cache := NewMultiCache()
	cache.Set("a", "1", len)
	for i := 0; i < b.N; i++ {
		cache.Print_in_mem()
	}
}
