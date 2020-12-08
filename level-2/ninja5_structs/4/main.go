package main

import "fmt"

func main() {
	anonStruc := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "jose",
		friends:   map[string]int{"jose": 1, "joana": 3},
		favDrinks: []string{"coca-cola", "juice"},
	}
	fmt.Println(anonStruc.first)
	fmt.Println(anonStruc.friends)
	fmt.Println(anonStruc.favDrinks)

	for k, v := range anonStruc.friends {
		fmt.Println(k, v)
	}

}
