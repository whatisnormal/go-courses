package main

import "fmt"

func main() {
	myFooInt := foo()
	myBarInt, myBarString := bar()
	fmt.Println("Hello", myFooInt, myBarInt, myBarString)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "hello"
}
