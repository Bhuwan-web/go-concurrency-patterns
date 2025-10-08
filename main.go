package main

import "fmt"

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
	// concurrencyrevisits.PoolReadmeFile("./README.md")
	// concurrencyrevisits.DisplaySimpleChannel()
	// concurrencyrevisits.DisplayChannelWithClose()
	// concurrencyrevisits.ScholarshipDistribution()
	// concurrencyrevisits.DisplayMultiSenderChannelStream()
	// concurrencyrevisits.DisplaySelectingSimultaneously()
	// patterns.DisplayLexicalConfinement()
	// patterns.DisplayAdHocConfinement()
	// patterns.DisplayForSelectWithChannel()
	// patterns.DisplayGoroutineLeak()
	// patterns.DisplayPreventingLeak()
	// patterns.DisplayOrChannelPattern()
	val := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", val[5:])
}
