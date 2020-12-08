package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   int
}

type ByAge []user

// Sort interface methods.

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	u1 := user{
		First: "James Bond",
		Age:   32,
	}
	u2 := user{
		First: "MoneyPenny",
		Age:   40,
	}

	u3 := user{
		First: "If The dog",
		Age:   2,
	}
	users := []user{u1, u2, u3}

	// Custom sort by Age interface
	sort.Sort(ByAge(users))

	fmt.Println(users)

}
