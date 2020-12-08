package main

import "fmt"

func main() {
	favSport := "surfing"

	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!!!")
	case "swimming":
		fmt.Println("go to the pool")
	case "surfing":
		fmt.Println("suringgggg")
	case "wingsuit flying":
		fmt.Println("WHAAAT")
	default:
		fmt.Println("Where to go?")
	}
}
