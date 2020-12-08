package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	flavors   []string
}

func main() {
	joana := person{
		firstName: "Joana",
		lastName:  "Roque",
		flavors:   []string{"baunilha", "chocolate"},
	}

	jose := person{
		firstName: "José",
		lastName:  "Fernandes",
		flavors:   []string{"ananas", "menta"},
	}

	fmt.Printf("%v %v gosta de: \n", joana.firstName, joana.lastName)
	for _, flavor := range joana.flavors {
		fmt.Println(flavor)
	}

	fmt.Printf("%v %v gosta de: \n", jose.firstName, jose.lastName)
	for _, flavor := range jose.flavors {
		fmt.Printf("\t%v", flavor)
	}
}
