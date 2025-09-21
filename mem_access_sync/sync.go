package mem_sync

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) performMemoryAccessSync() int {
	wg := sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			/*
				c.value++ is not an atomic operation, it is actually 3 operations
					1. read c.value
					2. increment c.value
					3. write c.value
				so, if 2 go routines are trying to increment c.value at the same time, they will read the same value, increment it and write it back, so the value will be incremented only by 1, not by 2
				so, we need to lock the value
			*/
			c.mu.Lock()
			defer c.mu.Unlock()
			c.value++
		}()
	}
	wg.Wait()
	return c.value
}
func DisplayMemorySync() {
	// this also doesn't grantee same number , but when a mutex locks a value, it won't throw a random number, either it will be zero or 100 at this point, either performMemoryAccessSync main thread get's the first number or waits for go routines to complete it's execution
	variation_map := make(map[int]int)
	for range 1000 {
		c := Counter{}
		val := c.performMemoryAccessSync()
		variation_map[val]++
	}
	fmt.Println(variation_map)
}

type Cache struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]string),
	}
}
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// c.mu.Lock()
	// defer c.mu.Unlock()
	val, ok := c.store[key]
	return val, ok
}
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}
func (c *Cache) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.store)
}
func DisplayCacheSync() {
	cache := NewCache()
	for i := range 5 {
		go func() {
			fmt.Printf("Write Goroutine %d : set value %d\n", i, i+1)
			cache.Set("key", fmt.Sprintf("value:%d", i+1))
		}()
	}

	for i := range 5 {
		go func() {
			for j := range 10 {
				value, _ := cache.Get("key")
				fmt.Printf("Read Goroutine %d: iteration %d : %s\n", i, j, value)
			}
			time.Sleep(time.Second)
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("cache count: ", cache.Count())
}
