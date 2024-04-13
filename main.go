package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	CACHE_SIZE  = 5
	dataSetSize = 1000_000_00
)

type Cache struct {
	LinkedList LinkedList
	Hash       Hash
}

func (c *Cache) Add(node *Node) {
	// Keep the refernce of current first value, which is also the right value of head.
	prevFirstValue := c.LinkedList.Head.Right

	// Set new first value of new node, by setting right node of head to it.
	c.LinkedList.Head.Right = node

	// Left of current node should point to head, and right of current node should point to prev first node.
	node.Left = c.LinkedList.Head
	node.Right = prevFirstValue
	prevFirstValue.Left = node

	c.LinkedList.Length += 1

	/* If we exceed size of the cache, we drop last element which
	is the least accessed element, so we consider this as one of cache invalidation rules */
	if c.LinkedList.Length > CACHE_SIZE {
		c.Remove(c.LinkedList.Tail.Left)
	}
}

func (c *Cache) Remove(node *Node) *Node {
	// Get the reference to current node left and right values nodes
	left := node.Left
	right := node.Right

	// Point previously fetched values to each other.
	left.Right = right
	right.Left = left

	// Remove provided node from cache hash, and decrement the total linked list length.
	delete(c.Hash, node.Value)
	c.LinkedList.Length -= 1

	return node
}

func (c *Cache) Check(n string) {
	var node *Node

	/* Check if value is in the cache hash; If it is, then remove it,
	   and add as recently used value; If not create and also add to cache hash. */
	if existingCacheValue, ok := c.Hash[n]; ok {
		node = c.Remove(existingCacheValue)
	} else {
		node = &Node{Value: n}
	}

	c.Add(node)
	c.Hash[n] = node
}

func (c *Cache) Display() {
	c.LinkedList.Display()
}

func (q *LinkedList) Display() {
	node := q.Head.Right

	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Hash map[string]*Node

func createCache() Cache {
	return Cache{
		LinkedList: createLinkedList(),
		Hash:       Hash{},
	}
}

func createLinkedList() LinkedList {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return LinkedList{
		Head:   head,
		Tail:   tail,
		Length: 0,
	}
}

func main() {
	elementsToCache := []string{"Terry", "Tee", "Dog", "Terry", "Car", "Terry"}

	cache := createCache()

	for _, e := range elementsToCache {
		cache.Check(e)
		cache.Display()
	}

	benchmarkLRUCache()
}

func benchmarkLRUCache() {
	cache := createCache()

	// Generate a large data set
	dataSet := generateLargeDataSet(dataSetSize)

	// Fill the cache with the large data set
	start := time.Now()
	for _, e := range dataSet {
		cache.Check(e)
	}
	fillElapsed := time.Since(start)

	// Measure the time taken to find one element in the cache hash
	randomIndex := rand.Intn(len(dataSet))
	searchElement := dataSet[randomIndex]

	start = time.Now()
	_, found := cache.Hash[searchElement]
	searchElapsed := time.Since(start)

	if !found {
		os.Exit(1)
	}

	fmt.Printf("Time taken to fill cache with %d elements: %s\n", dataSetSize, fillElapsed)
	fmt.Printf("Time taken to find element in cache hash: %s\n", searchElapsed)
}

// generateLargeDataSet generates a large data set for benchmarking purposes.
func generateLargeDataSet(size int) []string {
	dataSet := make([]string, size)
	for i := 0; i < size; i++ {
		dataSet[i] = fmt.Sprintf("Element%d", i)
	}
	return dataSet
}
