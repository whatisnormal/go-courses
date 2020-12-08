package main

import (
	"fmt"
)

func main() {
	f := returningInnerFunc()

	fmt.Println("Returning f: ", f())
}

func returningInnerFunc() func() int {

	return func() int {
		return 42
	}
}
