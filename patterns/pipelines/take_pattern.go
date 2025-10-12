package pipelines

import "fmt"

func take(done <-chan interface{}, intStream <-chan int, num int) <-chan int {
	result := make(chan int)
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
func DisplayTakePattern() {
	done := make(chan interface{})
	defer close(done)
	for i := range take(done, generator(done, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 5) {
		fmt.Println(i)
	}
}
