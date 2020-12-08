package main

import "fmt"

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}

	for i, v := range intArray {
		fmt.Println(i, v)

	}
	fmt.Printf("%T", intArray)
}
