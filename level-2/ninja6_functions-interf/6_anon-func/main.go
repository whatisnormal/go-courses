package main

import (
	"fmt"
)

func main() {
	func(i int) {
		fmt.Println("Printing from an anon func with value: ", i)
	}(3)
}
