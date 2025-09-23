package concurrencyrevisits

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// here if we remove the mutex, more than 1 goroutine can get the same value and increment it, leading to data race

	// assume two goroutine load 10, both will increment it to 11, but the final value will be 11, instead of 12

	// using mutex will prevent this, it locks the critical section from being accessed by more than one goroutine at a time
	c.value++
}

type SafeDeduct struct {
	balance int
	mu      sync.Mutex
}

func (sd *SafeDeduct) Deduct(deduct int) {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	if sd.balance >= deduct {
		// add intentional delay to allow threads to pick another goroutine to demonstrate race condition
		time.Sleep(time.Microsecond)
		sd.balance -= deduct
	}
}

func DisplayRaceCondition() {
	// race condition with no data race
	var wg sync.WaitGroup
	safeDeduct := SafeDeduct{balance: 100}
	deductBalance := func(deduct int) {
		safeDeduct.Deduct(deduct)
		wg.Done()
	}
	wg.Add(2)
	go deductBalance(70)
	go deductBalance(60)
	// for both the goroutines the total balance seems to be 100, while it pass through the condition, but will eventually reduce more than it have. It does't override the value thus generated, thus it's not a data race, but a race condition
	wg.Wait()
	if safeDeduct.balance < 0 {
		// print balance if race condition is present
		fmt.Println(safeDeduct.balance)
	}
}

func DisplayDataRace() {
	// race condition with data race
	iterations := 1000
	counter := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(counter.value)
}
