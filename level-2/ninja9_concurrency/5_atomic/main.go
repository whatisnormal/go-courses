package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var incrementer int64
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)

			loaderItem := atomic.LoadInt64(&incrementer)
			fmt.Println(loaderItem)
			wg.Done()

		}()

	}
	wg.Wait()

	// If you call this script several times
	// run 'go run -race main.go' to check race conditions detected.
	fmt.Println("End value: ", incrementer)
}
