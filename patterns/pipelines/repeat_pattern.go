package pipelines

import (
	"fmt"
	"time"
)

func repeat(done <-chan interface{}, values ...int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case result <- v:
				}
			}
		}
	}()
	return result
}

func DisplayRepeatPattern() {
	done := make(chan interface{})
	defer close(done)
	go func() {
		for i := range repeat(done, 1, 2, 3) {
			fmt.Println(i)
		}
	}()
	// repeat until 1 second of it's execution
	time.Sleep(1 * time.Second)
}
