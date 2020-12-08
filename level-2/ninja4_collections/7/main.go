package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken", "Not stirred"},
		{"Miss", "Moneypenny", "Hellooo james."},
	}

	for i, xs := range x {
		fmt.Println("Record: ", i)
		for _, val := range xs {
			fmt.Println(val)
		}
	}

}
