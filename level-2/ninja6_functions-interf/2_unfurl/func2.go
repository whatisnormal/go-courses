package main

import "fmt"

func main() {
	xi := []int{42, 44, 21, 44, 67, 88}
	total := foo(xi...)
	total2 := bar(xi)
	fmt.Printf("The sum of all values is: %v = %v \n", total, total2)
}

//This function cannot be called 'foo' as there is already another func called 'foo' in 'func1.go'
func foo(xi ...int) int {
	total := 0
	for _, el := range xi {
		total += el
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, el := range xi {
		total += el
	}
	return total
}
