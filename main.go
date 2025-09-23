package main

import concurrencyrevisits "githhub.com/bhuwan-web/learn-concurrency/concurrency_revisits"

// learning go concurrency in detail
func main() {
	// race.DisplayRace()
	// mem_sync.DisplayMemorySync()
	// waitgroups.DisplayWaitGroup()
	// mem_sync.DisplaySafeCounter()
	// mem_sync.DisplayCacheSync()
	// concurrencyrevisits.DisplayRaceCondition()
	// concurrencyrevisits.DisplayDataRace()
	// concurrencyrevisits.DisplayBroadcastingCondition()
	concurrencyrevisits.PoolReadmeFile("./README.md")
}
