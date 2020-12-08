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

	persons := map[string]person{
		joana.lastName: joana,
		jose.lastName:  jose,
	}

	fmt.Println(persons)

	for _, v := range persons {
		fmt.Println(v.firstName)

		for i, val := range v.flavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
