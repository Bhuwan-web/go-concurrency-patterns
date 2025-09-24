package main

import (
	"githhub.com/bhuwan-web/learn-concurrency/patterns"
)

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
	patterns.DisplayForSelectWithChannel()
}
