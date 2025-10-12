package patterns

import (
	"fmt"
)

// try handling exceptions in main function rather than on goroutines itself, as it would have more context of the program,
type Result struct {
	Response interface{}
	Error    error
}

func callWebsite(name string) (string, error) {
	// mimic network call

	if name == "error" {
		return "", fmt.Errorf("something went wrong")
	}
	return "Successfully fetched data", nil
}
func getResults(done <-chan interface{}, urls ...string) <-chan Result {
	response := make(chan Result)
	go func() {
		defer close(response)
		for _, url := range urls {
			nw_response, err := callWebsite(url)
			select {
			case <-done:
				return
			case response <- Result{Response: nw_response, Error: err}:
			}
		}
	}()
	return response
}
func DisplayExceptionHandling() {
	done := make(chan interface{})
	defer close(done)
	error_count := 0
	error_threshold := 3
	urls := []string{"apple", "facebook", "error", "error", "error", "google", "amazon"}
	for result := range getResults(done, urls...) {
		if result.Error != nil {
			// exception handling is being here in the main function, as it has more context of the program, like error threshold and error counts
			fmt.Println("Error: ", result.Error)
			error_count++
			if error_count >= error_threshold {
				fmt.Println("Error threshold reached, breaking the program")
				break
			}
			continue
		}
		fmt.Println("Result: ", result.Response)
	}
}
