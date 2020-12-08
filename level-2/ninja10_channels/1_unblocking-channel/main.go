package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// we could use also buffered channel by declaring like: c := make(chan int, 1) but is not the ideal solution

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
