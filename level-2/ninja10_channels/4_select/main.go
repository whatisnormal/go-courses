package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q <-chan int) {
	for {
		select {
		// is there anything on 'c'?
		case v := <-c:
			fmt.Println(v)
			// is there anything on 'q'? If so, return.
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		q <- 1
		close(c)
	}()

	return c
}
