package patterns

import (
	"fmt"
	"time"
)

func DisplayGoroutineLeak() {
	helloFunc := func(strings <-chan string) <-chan struct{} {
		result := make(chan struct{})
		go func() {
			defer close(result)
			defer fmt.Println("Goroutine leaks. If this goroutine was ever closed, this would have been triggered and shown to the terminal")
			for str := range strings {
				// here since strings is passed as nil value, it keeps on waiting forever and will never get the value from the channel and will keep on running.
				// goroutines are not garbage collected, they keep on running and will keep on using the memory.
				fmt.Println("hello", str)
			}
		}()
		return result
	}

	// ? uncommenting this section prevents goroutine leaks itself. As there are three methods to prevent leaks
	// 1. complete it's work
	// 2. error
	// 3. when it's told to stop working

	helloStream := helloFunc(nil)
	time.Sleep(10 * time.Second)
	<-helloStream
}
func DisplayPreventingLeak() {
	// prevent goroutine leaks
	// use done channel to signal the goroutine to stop working
	done := make(chan struct{})
	helloFunc := func(done <-chan struct{}, strings <-chan string) <-chan struct{} {
		// convention to pass done channel as first argument
		result := make(chan struct{})
		go func() {
			// this pattern of returning channel as seen in confinement pattern
			defer close(result)
			defer fmt.Println("Goroutine leaks. If this goroutine was ever closed, this would have been triggered and shown to the terminal")
			for {
				select {
				case <-done:
					// signal to stop working
					return
				case str := <-strings:
					fmt.Println("hello", str)
				}
			}
		}()
		return result
	}

	// ? uncommenting this section prevents goroutine leaks itself. As there are three methods to prevent leaks
	// 1. complete it's work
	// 2. error
	// 3. when it's told to stop working

	result := helloFunc(done, nil)
	time.Sleep(1 * time.Second)
	close(done) //done channel is closed to signal the goroutine to stop working
	<-result
}
