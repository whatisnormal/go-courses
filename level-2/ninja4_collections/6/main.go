package main

import "fmt"

func main() {
	x := make([]string, 2, 3)
	x = []string{`Alabama`, `Alaska`, `Arizona`}

	fmt.Println(x)

	fmt.Println(len(x))

	fmt.Println(cap(x))

}
