package main

import "fmt"

func main() {
	// Create a map - Method 1
	/* 	colors := map[string]string{
	   		"red":   "#ff0000",
	   		"green": "#4bf745",
	   	}
	   	fmt.Println(colors) */

	//***********

	// Create a map - Method 2
	//var colors map[string]string

	//***********

	// Create a map - Method 3
	/* 	colors := make(map[int]string)
	   	colors[10] = "#ffffff"
	   	//Delete the record whith key '10'
	   	delete(colors, 10)
		   fmt.Println(colors) */

	//***********

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
