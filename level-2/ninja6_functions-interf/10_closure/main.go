package main

import (
	"fmt"
)

//Closure is when you are enclosing the scope of a variable.
func main() {
	f := foo()
	// By calling 'f' several times, it is expected that 'x' will increment
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := foo()

	// by re-scoping / initializing, 'g' will start with 'x' with its initial value.

	fmt.Println(g())
	fmt.Println(g())

}

func foo() func() int {
	//'x' is enclosed
	x := 0
	return func() int {
		x++
		return x
	}

}
