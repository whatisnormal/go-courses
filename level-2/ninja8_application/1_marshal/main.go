package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
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

	fmt.Println(users)

	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
