package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Puts 100 numbers into a channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// Pulls the numbers and print
	// range can only read after the channel is closed.
	for val := range c {
		fmt.Println(val)
	}

}
