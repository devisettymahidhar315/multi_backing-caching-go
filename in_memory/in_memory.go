package in_memory

import (
	"container/list"
	"fmt"
)

type CacheNode struct {
	key   string
	value string
}

type LRUCache struct {
	cache map[string]*list.Element
	list  *list.List
}

func NewLRUCache() *LRUCache {
	return &LRUCache{

		cache: make(map[string]*list.Element),
		list:  list.New(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {

	if elem, found := c.cache[key]; found {
		c.list.MoveToFront(elem)
		return elem.Value.(*CacheNode).value, true
	}
	return "", false
}

func (c *LRUCache) Put(key string, value string, maxLength int) {

	if elem, found := c.cache[key]; found {

		elem.Value.(*CacheNode).value = value
		c.list.MoveToFront(elem)
		return
	}
	if c.list.Len() == maxLength {

		evicted := c.list.Back()
		if evicted != nil {
			c.list.Remove(evicted)
			delete(c.cache, evicted.Value.(*CacheNode).key)

		}
	}

	newNode := &CacheNode{key: key, value: value}
	entry := c.list.PushFront(newNode)
	c.cache[key] = entry
}

func (c *LRUCache) Print() {
	fmt.Println("in memeory")

	for elem := c.list.Front(); elem != nil; elem = elem.Next() {
		node := elem.Value.(*CacheNode)
		fmt.Printf("%s: %s\n", node.key, node.value)
	}
}

func (c *LRUCache) Del(key string) {
	if elem, found := c.cache[key]; found {
		c.list.Remove(elem)
		delete(c.cache, elem.Value.(*CacheNode).key)
	} else {
		fmt.Println("not found")
	}
}
