package main

import (
	"fmt"
)

type LRUCache struct {
	Capacity int
	Cache    map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Cache:    map[int]int{},
	}
}

func (l *LRUCache) Get(key int) int {
	return 0
}

func (l *LRUCache) Put(key, value int) {
}

func main() {
	c := Constructor(5)
	fmt.Printf("cache lenght: %d\n", c.Capacity)
}
