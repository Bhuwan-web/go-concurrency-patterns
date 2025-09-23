package concurrencyrevisits

import "fmt"

func DisplayMultiSenderChannelStream() {
	intStream := make(chan int)
	go func() {
		for i := range 10 {
			intStream <- i
		}
		close(intStream)
	}()
	for value := range intStream {
		fmt.Printf("Received value %d\n", value)
	}
}
