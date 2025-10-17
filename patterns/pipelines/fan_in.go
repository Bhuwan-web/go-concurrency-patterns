package pipelines

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	// this function ensures that even if multiple channels are created, it takes it all and return one by one, and follow the function definition of pipeline pattern used here
	var wg sync.WaitGroup
	multiplexStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(multiplexStream)
	}()
	return multiplexStream
}

func takeI(done <-chan interface{}, intStream <-chan interface{}, num int) <-chan interface{} {
	result := make(chan interface{})
	go func() {
		defer close(result)
		for range num {
			select {
			case <-done:
				return
			case result <- <-intStream:
			}
		}
	}()
	return result
}
func DisplayFanInPattern() {
	done := make(chan interface{})
	defer close(done)

	num_cpus := runtime.NumCPU()
	randInt := func() interface{} {
		int_val := rand.Intn(5000000000)
		return int_val
	}
	randIntStream := toInt(done, repeatFn(done, randInt))
	finders := make([]<-chan interface{}, num_cpus)
	start := time.Now()

	for i := 0; i < num_cpus; i++ {
		// create Multiple primeFinder function call, where each primeFinder function call is a goroutine and return a channel of prime numbers, each one computes prime numbers in parallel and returns a channel of prime numbers, not a prime value, this means, even a single one has a capacity to find prime sequentially but it uses multiple ones to find prime numbers in parallel
		finders[i] = primeFinder(done, randIntStream)
	}
	for i := range takeI(done, fanIn(done, finders...), 10) {
		fmt.Println(i)
	}
	fmt.Printf("%v seconds elapsed\n", time.Since(start).Seconds())
}
