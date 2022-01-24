package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mtx      sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if val, ok := c.items[key]; ok {
		val.Value.(*cacheItem).value = value
		c.queue.MoveToFront(val)
		return true
	}
	if c.capacity == c.queue.Len() {
		last := c.queue.Back()
		c.queue.Remove(last)
		delete(c.items, last.Value.(*cacheItem).key)
	}
	c.items[key] = c.queue.PushFront(&cacheItem{key, value})
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if val, ok := c.items[key]; ok {
		c.queue.MoveToFront(val)
		return val.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for i := c.queue.Back(); i != nil; {
		c.queue.Remove(i)
		delete(c.items, i.Value.(*cacheItem).key)
	}
}
