package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//Created muted
	var mutex sync.Mutex

	incrementer := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			//Lock mutex
			mutex.Lock()
			v := incrementer

			//Forces scheduler to jump to other go routine.
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()

			//Unlock mutex
			mutex.Unlock()
		}()

	}
	wg.Wait()

	// If you call this script several times
	// run 'go run -race main.go' to check race conditions detected.
	fmt.Println("End value: ", incrementer)
}
