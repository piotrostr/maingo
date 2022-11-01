package main

import (
	"log"
	"os"
	"sync"
)

func process(val string) {}

var logger = log.New(os.Stdout, "", log.LstdFlags)

func main() {
	values := []string{"a", "b", "c"}
	wg := sync.WaitGroup{}
	wg.Add(len(values))
	for _, val := range values {
		go func(val string) {
			defer wg.Done()
			logger.Printf("Processing %s...", val)
			process(val)
		}(val)
	}
	wg.Wait()
	logger.Println("Done")
}
