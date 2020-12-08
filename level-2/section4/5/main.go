package main

import "fmt"

type myType int

var x myType

// var 'y' is of type 'x'
var y int

func main() {
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 3

	fmt.Println(x)

	y = int(x)

	fmt.Println(y)

	fmt.Printf("%T\n", y)

}
