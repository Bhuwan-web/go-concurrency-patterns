package patterns

import "fmt"

// 1. ad hoc confinement
//  a set of rules that are followed by the programmer, following the convention

func DisplayAdHocConfinement() {
	// confinement is not a feature of go, it's a convention
	data := make([]int, 10)
	// here, chan can be updated or written from anywhere in this function
	// but by convention, it's confined to be used only in the readNumbers function
	// it is easily breakable and hard to regulate
	readNumbers := func(readStream chan int) {
		defer close(readStream)
		for i := range data {
			readStream <- i
		}
	}
	readStream := make(chan int)
	go readNumbers(readStream)
	for i := range readStream {
		fmt.Println("ad hoc confinement", i)
	}
}

// 2. lexical confinement

// confinement via ownership, code such that the data can never be updated outside the lexical scope of that particular function, so the data is confined to that function

func DisplayLexicalConfinement() {
	data := make([]int, 10)
	readNumbers := func() <-chan int {
		// here, data can't be updated or written from anywhere in this channel as it returns a read-only channel
		// it is confined to this function
		// the channel can only be read outside of this function, this function owns that channel, create it, close it and write to it.
		readStream := make(chan int)
		go func() {
			defer close(readStream)
			for i := range data {
				readStream <- i
			}
		}()
		return readStream
	}

	readStream := readNumbers()
	for i := range readStream {
		fmt.Println("lexical confinement", i)
	}
}
