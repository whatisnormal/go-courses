package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c) // close sends a message to the channel with '0, false'

	v, ok = <-c
	fmt.Println(v, ok)
}
