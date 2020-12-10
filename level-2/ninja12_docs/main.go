package main

import (
	"log"

	dogs_utils "example.com/username/repo/ninja12_docs/dog"
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
		age:  dogs_utils.ConvertToHumanYears(humanYears),
	}

	log.Println(ifTheDog)
}
