package pipelines

import (
	"fmt"
	"math/rand"
	"time"
)

// this is a simple prime number finder function, without fanin pattern implemented, to demonstrate the pipeline pattern and it's performance
func isPrime(i int) bool {
	// here we check if a number is prime or not
	for j := i - 1; j > 2; j-- {
		if i%j == 0 {
			return false
		}
	}
	return true
}
func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primes := make(chan interface{})
	go func() {
		defer close(primes)
		for i := range intStream {
			// we try to find prime numbers, forever loop until it gets a prime number
			if isPrime(i) {
				select {
				case <-done:
					return
				case primes <- i:
					// we found a prime number
				}
			}
			// else we continue to find prime numbers
		}
	}()
	return primes
}
func toInt(done <-chan interface{}, intStream <-chan interface{}) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := range intStream {
			select {
			case <-done:
				return
			default:
				result <- i.(int)
			}
		}
	}()
	return result
}
func DisplayNonFanInPattern() {
	done := make(chan interface{})
	defer close(done)

	randInt := func() interface{} {
		int_val := rand.Intn(5000000000)
		return int_val
	}
	// here we generate random numbers and try to find prime numbers
	start := time.Now()
	// this is a helper functions but not part of pipelines directly as it doesn't follow same function definition as other functions
	randIntStream := toInt(done, repeatFn(done, randInt))
	// created pipeline to find prime numbers

	// pipeline function definition
	// interface{}, <-chan int returns <-chan interface{}

	pipeline := takeI(done, primeFinder(done, randIntStream), 10)
	for i := range pipeline {
		fmt.Println(i)
	}
	fmt.Printf("%v seconds elapsed\n", time.Since(start).Seconds())
}
