package main

import (
	"fmt"
)

func main() {
	//cs := make(chan<- int) // this channel is of type 'send' but there is an instruction below in line 13 which receives. Solution is to change to a general channel (bidirect)
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

}
