package main

import (
	dog "example.com/username/repo/ninja13_tests-n-benchmarks/1_basics/dog"
	"log"
)

type canine struct {
	name string
	age  int
}

// run docs server with -> godoc -http=0.0.0.0:8080
func main() {

	humanYears := 10
	ifTheDog := canine{
		name: "If",
		age:  dog.Years(humanYears),
	}

	log.Println(ifTheDog)

	log.Println(dog.YearsTwo(2))
}
