package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := createCache()

	// Test adding elements
	elementsToCache := []string{"Dog", "Cat", "Soda", "Tee", "Dog", "Terry", "Car"}
	for _, e := range elementsToCache {
		cache.Check(e)
	}

	// Test cache behavior
	expectedCacheState := []string{"Car", "Terry", "Dog", "Tee", "Soda"}
	actualCacheState := getCacheState(cache)
	if !equalSlice(expectedCacheState, actualCacheState) {
		t.Errorf("Expected cache state: %v, but got: %v", expectedCacheState, actualCacheState)
	}

	// Test adding more elements than cache size
	elementsToCache = []string{"Apple", "Banana", "Grape", "Pineapple", "Watermelon"}
	for _, e := range elementsToCache {
		cache.Check(e)
	}

	// Test cache state after adding more elements than cache size
	expectedCacheState = []string{"Watermelon", "Pineapple", "Grape", "Banana", "Apple"}
	actualCacheState = getCacheState(cache)
	if !equalSlice(expectedCacheState, actualCacheState) {
		t.Errorf("Expected cache state: %v, but got: %v", expectedCacheState, actualCacheState)
	}
}

func getCacheState(cache Cache) []string {
	state := make([]string, 0, CACHE_SIZE)
	node := cache.LinkedList.Head.Right
	for i := 0; i < cache.LinkedList.Length; i++ {
		state = append(state, node.Value)
		node = node.Right
	}
	return state
}

func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
