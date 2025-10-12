package pipelines

import "fmt"

func generator(done <-chan interface{}, integers []int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := range integers {
			select {
			case <-done:
				return
			case result <- i:
			}
		}
	}()
	return result
}
func addStream(done <-chan interface{}, intStream <-chan int, adder int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := range intStream {
			select {
			case <-done:
				return
			case result <- i + adder:
			}
		}
	}()
	return result
}
func multipleStream(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := range intStream {
			select {
			case <-done:
				return
			case result <- i * multiplier:
			}
		}
	}()
	return result
}
func DisplayChannelPipeline() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// here it removes the issue of increasing function calls of stream processing and cluttering memory footprints of batch processing as it now uses channels
	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, ints)
	pipeline := addStream(done, multipleStream(done, addStream(done, intStream, 1), 2), 3)

	for i := range pipeline {
		fmt.Println(i)
	}
}
