package main

import "fmt"

func main() {
	sliceOfInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 19}

	for i := range sliceOfInts {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
