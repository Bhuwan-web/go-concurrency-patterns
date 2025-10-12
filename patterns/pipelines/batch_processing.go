package pipelines

import "fmt"

//	base properties of pipeline,
//
// 1. it is a sequence of stages
// 2. each stage takes a same datatype as input and response
// 3. each stage is a function that takes a slice of data and returns a slice of data
// 4. each stage is a function that takes a slice of data and returns a slice of data
// 5. a function must be reified by the language so that it may be passed around
func add(ints []int, adder int) []int {
	for i := range ints {
		ints[i] += adder
	}
	return ints
}
func multiple(ints []int, multiplier int) []int {
	for i := range ints {
		ints[i] *= multiplier
	}
	return ints
}
func DisplayBatchProcessing() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// creating batch processing pipeline , such that it process all data in back at a time. Meaning all the array will be added first and then transformed array will then be multiplied.
	// it leaves the memory footprint to x2 for each stage of pipeline
	for i := range multiple(add(ints, 1), 2) {
		fmt.Println("Batch processing", i)
	}
}
