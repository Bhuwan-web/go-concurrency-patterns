package pipelines

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	result := make(chan interface{})
	go func() {
		defer close(result)
		for {
			select {
			case <-done:
				return
			case result <- fn():
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

func DisplayRepeatFnPattern() {
	// dynamic random numbers generation until 1 second of it's execution

	done := make(chan interface{})
	defer close(done)
	rand := func() interface{} { v, _ := rand.Int(rand.Reader, big.NewInt(1000)); return v }

	go func() {
		for i := range repeatFn(done, rand) {
			fmt.Println(i)
		}
	}()
	// repeat until 1 second of it's execution
	time.Sleep(1 * time.Second)
}
