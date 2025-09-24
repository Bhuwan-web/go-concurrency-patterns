package patterns

import (
	"fmt"
	"time"
)

func DisplayForSelect() {
	for {
		select {
		default:
			fmt.Println("default")
		}
		break
	}
}

func DisplayForSelectWithChannel() {
	channel := make(chan int)
	done := make(chan struct{})
	iteration := 0
	go func() {
		for {
			select {
			case <-channel:
				fmt.Println("received channel message ")
				iteration++
				done <- struct{}{}
				return
			default:
				fmt.Println("default")
			}
			iteration++
			fmt.Println("iteration: ", iteration)
		}
	}()
	time.Sleep(4 * time.Microsecond)
	channel <- 1
	<-done
}
