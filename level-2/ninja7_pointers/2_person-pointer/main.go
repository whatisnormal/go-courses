package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Miss Moneypenny" // short hand to dereferencing -> (*p).
	//same as above -> (*p).name = "Miss Moneyp"
}

func main() {
	p1 := person{
		name: "James Bond",
	}

	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)

}
