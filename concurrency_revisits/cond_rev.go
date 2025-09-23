package concurrencyrevisits

import (
	"fmt"
	"sync"
	"time"
)

var c = sync.NewCond(&sync.Mutex{})
var queue = make([]int, 0, 10)

func dequeue() {
	time.Sleep(time.Second)
	c.L.Lock()
	defer c.L.Unlock()
	queue = queue[1:]
	fmt.Println("Removed one element from queue")
	c.Signal()
}

func enqueue(item int) {
	c.L.Lock()
	defer c.L.Unlock()
	for len(queue) == 2 {
		fmt.Println("Queue is full, waiting for space to be available")
		c.Wait()
	}
	queue = append(queue, item)
	fmt.Println("Added one element to queue")
}

func DisplayConditionVariable() {
	// takes 10 queue values, but 2 elements at a time
	for i := 0; i < 10; i++ {
		enqueue(i)
		go dequeue()
	}
}
