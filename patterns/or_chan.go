package patterns

import (
	"fmt"
	"time"
)

// or function takes a variadic number of channels and returns a single channel
// it returns a channel that will be closed when any of the input channels are closed
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			// optimization for two channels, as for this recursion case, it should at least have two elements, and reduce the number of goroutines created
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
				// special case in go, if index 3 is out of range, it still will return empty array
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
func runAfter(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
		fmt.Println("runAfter", after)
	}()
	return c
}
func DisplayOrChannelPattern() {
	value := or(runAfter(3*time.Second), runAfter(4*time.Second))
	<-value
}
