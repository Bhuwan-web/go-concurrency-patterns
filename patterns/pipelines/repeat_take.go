package pipelines

import "fmt"

func DisplayRepeatTake() {
	done := make(chan interface{})
	defer close(done)
	// combination of take and repeat pattern
	for i := range take(done, repeat(done, 1, 2), 5) {
		fmt.Println(i)
	}
}
