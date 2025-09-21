package mem_sync

import "sync"

type safeCounter struct {
	mu    sync.RWMutex
	value int
}

func (c *safeCounter) PerformSafeCounter() {
	c.value = 100
	for range 1000 {
		go func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			c.value++
		}()
	}
	for i := range 1000 {
		go func() {
			c.mu.RLock()
			defer c.mu.RUnlock()
			println(i, c.value)
		}()
	}
}

func DisplaySafeCounter() {
	counter := safeCounter{}
	counter.PerformSafeCounter()
}
