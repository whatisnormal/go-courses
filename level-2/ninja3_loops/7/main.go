package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	} else if x == "Moneypenny" {
		fmt.Println(x)
	} else {
		fmt.Println("Neither")
	}
}
