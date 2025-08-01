package mem_sync

import (
	"fmt"
	"sync"
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
