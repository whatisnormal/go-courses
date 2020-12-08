package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
	fmt.Println("")

	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
