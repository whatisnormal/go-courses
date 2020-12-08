package main

import (
	"fmt"
)

func main() {
	f := func(i int) {
		fmt.Println("Printing from an anon func with value: ", i)
	}

	f(2)
}
