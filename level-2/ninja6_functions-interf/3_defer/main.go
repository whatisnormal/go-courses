package main

import "fmt"

func main() {
	xi := []int{42, 44, 21, 44, 67, 88}
	defer printAll(xi...)
	printMe()
}

//This function cannot be called 'foo' as there is already another func called 'foo' in 'func1.go'
func printAll(xi ...int) {
	total := 0
	for _, el := range xi {
		total += el
	}
	fmt.Printf("Was I printed second with defer? The sum of all values is: %v\n", total)

}

func printMe() {
	fmt.Println("Was I printed first?")
}
