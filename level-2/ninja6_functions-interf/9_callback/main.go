package main

import (
	"fmt"
)

func main() {
	result := callbackTest(func() int {
		return 42
	})

	fmt.Println("Returning result: ", result)
}

func callbackTest(f func() int) int {
	result := f()
	return result + 1

}
