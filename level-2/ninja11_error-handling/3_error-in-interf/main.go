package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", c.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran", e, "\n")
}
