package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	//numRoutines := 10

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {

			c <- 1 // this is  blocking operation. If no one reads it, it will block the goroutine
			fmt.Println("Sent ", 1)
			wg.Done()
			close(c)
		}()
	}

	fmt.Println("About to wait")

	go func() {
		var sum int
		for i := range c {
			sum += i
		}

		fmt.Println("Sum is: ", sum)
	}()

	wg.Wait()
	close(c)
	fmt.Println("All go routines finished")

}
