package main

import "fmt"

func main() {
	intArray := []int{1, 2, 3, 4, 5, 33, 44, 555, 633, 22}

	for i, v := range intArray {
		fmt.Println(i, v)

	}
	fmt.Printf("%T", intArray)
}
