package redis

import (
	"fmt"
	"testing"
)

// Length of the cache for testing
const len1 = 2

// TestRedisPut tests the Put method of the Redis cache
func TestRedisPut(t *testing.T) {
	// Initialize a new Redis cache
	cache := NewLRUCache()

	// Insert key-value pairs into the cache
	cache.Put("a1", "1", len1)
	cache.Put("b1", "2", len1)

	// Retrieve the value for key "a1"
	result := cache.Get("a1")

	// Check if the value is as expected
	if result != "1" {
		t.Error("Expected value '1', got", result)
	}
}

// TestRedisGet tests the Get method of the Redis cache
func TestRedisGet(t *testing.T) {
	// Initialize a new Redis cache
	cache := NewLRUCache()

	// Insert key-value pairs into the cache
	cache.Put("a1", "1", len1)
	cache.Put("b1", "2", len1)

	// Retrieve the value for key "a1"
	result := cache.Get("a1")
	if result != "1" {
		t.Error("Expected value '1', got", result)
	}

	// Attempt to retrieve a value for a non-existent key "v"
	result = cache.Get("v")
	if result != "" {
		t.Error("Expected empty string for non-existent key, got", result)
	}
}

// TestRedisPrint tests the Print method of the Redis cache
func TestRedisPrint(t *testing.T) {
	// Initialize a new Redis cache
	cache := NewLRUCache()

	// Insert key-value pairs into the cache
	cache.Put("a1", "1", len1)
	cache.Put("b1", "2", len1)

	// Print the current state of the cache
	result := cache.Print()
	fmt.Println(result)
	expected_result := "b1:2, a1:1"

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}

	// Insert another key-value pair to exceed the cache length
	cache.Put("c1", "3", len1)

	// Print the current state of the cache
	result = cache.Print()
	expected_result = "c1:3, b1:2"

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}
}

// TestRedisDel tests the Del method of the Redis cache
func TestRedisDel(t *testing.T) {
	// Initialize a new Redis cache
	cache := NewLRUCache()

	// Insert a key-value pair into the cache
	cache.Put("a1", "1", len1)

	// Delete the key "a1"
	cache.Del("a1")

	// Print the current state of the cache
	result := cache.Print()
	expected_result := ""

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}

	// Insert key-value pairs into the cache
	cache.Put("a1", "1", len1)
	cache.Put("b1", "2", len1)

	// Delete the key "a1"
	cache.Del("a1")

	// Attempt to retrieve the value for the deleted key "a1"
	result = cache.Get("a1")
	expected_result = ""

	// Check if the result matches the expected result
	if result != expected_result {
		t.Error("Expected empty string for deleted key, got", result)
	}
}
