package concurrencyrevisits

import (
	"fmt"
	"os"
	"sync"
)

var newPool = &sync.Pool{
	New: func() interface{} {
		fmt.Println("Creating new instance")
		return struct{}{}
	},
}

func DisplayPoolingConcept() {
	newPool.Get()             // creating new instance
	instance := newPool.Get() // creating another new instance
	newPool.Put(instance)     // putting instance back to pool
	newPool.Get()             // reusing instance
}

func PoolReadmeFile(filePath string) {
	// pooling memory allocation for file reading, defined size of memory. Read file in chunks of 1024 bytes, 4 chunks at a time
	var iteration int
	fileReader := sync.Pool{
		New: func() interface{} {
			iteration++
			mem := make([]byte, 1024)
			return &mem
		},
	}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileReader.Put(fileReader.New())
	fileReader.Put(fileReader.New())
	fileReader.Put(fileReader.New())
	fileReader.Put(fileReader.New())
	defer file.Close()

	for {
		buf := fileReader.Get().(*[]byte)
		n, err := file.Read(*buf)
		defer fileReader.Put(buf)
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		fmt.Printf("%s", (*buf)[:n])
	}
}
