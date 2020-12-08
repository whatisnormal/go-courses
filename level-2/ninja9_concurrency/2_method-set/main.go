package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello ", p.first)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{first: "Jose"}

	// saySomething(p1) // does not work
	saySomething(&p1)

	p1.speak()
}
