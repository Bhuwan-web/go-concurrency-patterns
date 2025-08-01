package race

import (
	"fmt"
	"time"
)

func performRace() int {
	var num int
	go func() {
		for i := range 100 {
			num = i
		}
	}()
	time.Sleep(1 * time.Microsecond)
	return num
}

func DisplayRace() {
	fmt.Println("Displaying race condition on executing multiple executions ")
	seen_count := make(map[int]int)
	fmt.Println("Performing same functions multiple time")
	for range 2000 {
		val := performRace()
		seen_count[val]++
	}
	fmt.Println("Number of unique values generated")
	fmt.Print(seen_count)
}
