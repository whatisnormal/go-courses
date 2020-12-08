package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	myTruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "orange",
		},
		luxury: false,
	}

	fmt.Println(myTruck)
	fmt.Println(mySedan)

	fmt.Println(myTruck.color)

	fmt.Println(mySedan.doors)
}
