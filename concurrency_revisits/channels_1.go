package concurrencyrevisits

import (
	"fmt"
	"sync"
	"time"
)

func DisplaySimpleChannel() {
	ch := make(chan int)
	go func() {
		// send value to channel
		ch <- 1
	}()
	// receive value from channel
	fmt.Println(<-ch)
}

func DisplayChannelWithClose() {
	//  this code block is expected to flow one data through channel and close it for the rest of the receivers.
	//  it first makes sure one value got the value and signal close once the values are well fed
	ch := make(chan int)
	putFirstValue := sync.NewCond(&sync.Mutex{})
	go func() {
		// send value to channel once
		ch <- 1
		time.Sleep(time.Second)
		putFirstValue.Signal()
	}()
	// close channel before receiving value from channel
	go func() {
		putFirstValue.L.Lock()
		putFirstValue.Wait()
		defer putFirstValue.L.Unlock()
		// channel only stream values once, so it will be closed after first iteration and all other iterations will get the default values from channel
		close(ch)
	}()
	for range 10 {
		// receive values from channels multiple times
		fmt.Println(<-ch)
	}
}

func ScholarshipDistribution() {
	scholarshipStream := make(chan int)
	ScholarshipCount := 10
	applicantCount := 100
	wg := sync.WaitGroup{}
	for i := range applicantCount {
		wg.Add(1)
		go func(applicant int) {
			defer wg.Done()
			_, ok := <-scholarshipStream
			if !ok {
				// fmt.Println("No more scholarships available")
				return
			}
			fmt.Printf("Applicant %d got scholarship\n", applicant)
		}(i)
	}
	for range ScholarshipCount {
		// send value to channel
		scholarshipStream <- 1
		// fmt.Println("Scholarship distributed")
	}
	close(scholarshipStream)
	wg.Wait()
}
