package concurrencyrevisits

import (
	"fmt"
	"sync"
	"time"
)

var message_cond = sync.NewCond(&sync.Mutex{})

func sendMessage(fn func()) {
	messageWg := sync.WaitGroup{}
	messageWg.Add(1)
	go func() {
		message_cond.L.Lock()
		messageWg.Done()
		defer message_cond.L.Unlock()
		message_cond.Wait()
		fn()
	}()
	messageWg.Wait()
}

func DisplayBroadcastingCondition() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	sendMessage(func() { fmt.Println("Sending message via Email"); wg.Done() })
	sendMessage(func() { fmt.Println("Sending message via SMS"); wg.Done() })
	sendMessage(func() { fmt.Println("Sending message via WhatsApp"); wg.Done() })
	time.Sleep(1 * time.Millisecond) //mimic some work before sending message
	message_cond.Broadcast()
	wg.Wait()
}
