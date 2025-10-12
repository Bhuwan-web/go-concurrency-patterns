package pipelines

import "fmt"

func add_(num, adder int) int {
	return num + adder
}
func multiple_(num, multiplier int) int {
	return num * multiplier
}
func DisplayStreamProcessing() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range ints {
		// here one by one element is processed and transformed limiting memory access.
		// but at a cost of increasing function calls
		fmt.Println("Stream processing", multiple_(add_(multiple_(i, 2), 1), 2))
	}
}
