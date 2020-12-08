package main

import "fmt"

type person struct {
	firsName string
	lastName string
	age      int
}

func main() {
	p1 := person{
		firsName: "jose",
		lastName: "fernandes",
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Printf("This guy %v with last name %v has spoken.", p.firsName, p.lastName)
}
