package waitgroups

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) PerformWithoutWaitGroup() int {
	go func() {
		for range 1000 {
			c.value++
		}
	}()
	// executes this line without waiting for goroutines to complete, between the goroutines iterations, or even before goroutines get chance to get started
	return c.value
}
func (c *Counter) PerformWithWaitGroup() int {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 1000 {
			c.value++
		}
	}()
	wg.Wait()
	return c.value
}

func DisplayWaitGroup() {
	var c1 Counter
	var c2 Counter

	fmt.Printf("Without WaitGroup: %d\n", c1.PerformWithoutWaitGroup())
	fmt.Printf("With wg: %d\n", c2.PerformWithWaitGroup())
}
