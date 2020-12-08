package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Create a struct - Method 1
	/* 	alex := person{firstName: "Alex",
		lastName: "Anderson"}
	fmt.Println(alex) */

	// Create a struct - Method 2
	/* 	var alex person
	   	alex.firstName = "Alex"
	   	alex.lastName = "Anderson"
	   	fmt.Println(alex)
	   	fmt.Println("%+v", alex) */

	//******** Example of a embeded struct inside another

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//fmt.Println("%+v", jim)

	// & refers to the memory address that jim resides
	/* 	jimPointer := &jim
	   	jimPointer.updateName("jimmy")
		   jim.print() */

	jim.updateName("jimmy")
	jim.print()
}

//*person = pointer to a person.
func (pointerToPerson *person) updateName(newFirstName string) {
	//*pointerToPerson = The value that the pointer is pointing to.
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println("%+v", p)
}
