package concurrencyrevisits

import "fmt"

func DisplaySelecting() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go func() {
		for i := range 10 {
			channel1 <- i
		}
		close(channel1)
	}()
	go func() {
		for i := range 10 {
			channel2 <- i
		}
		close(channel2)
	}()
	for i := range 20 {
		select {
		case value := <-channel1:
			fmt.Println(i, "Received value from channel 1", value)
		case value := <-channel2:
			fmt.Println(i, "Received value from channel 2", value)
		}
	}
}

func DisplaySelectingSimultaneously() {
	channel1 := make(chan int)
	close(channel1)
	channel2 := make(chan int)
	close(channel2)
	chan1count := 0
	chan2count := 0

	for range 1000 {
		select {
		// if both channels are ready, it will select one at random
		case <-channel1:
			chan1count++
		case <-channel2:
			chan2count++
		}
	}
	fmt.Println("Channel 1 count", chan1count)
	fmt.Println("Channel 2 count", chan2count)
}
