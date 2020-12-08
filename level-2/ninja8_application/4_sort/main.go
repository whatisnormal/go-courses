package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   int
}

func main() {
	xi := []int{2, 5, 3, 6, 7, 2, 6, 8, 8, 9, 2}
	xs := []string{"joana", "pedro", "miguel", "laureano", "paula", "ana"}

	fmt.Println(xi)

	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
