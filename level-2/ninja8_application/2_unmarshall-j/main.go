package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

type user2 struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
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
	jsonInString := string(bs)
	fmt.Println(jsonInString)

	var otherUsers []user2

	err2 := json.Unmarshal(bs, &otherUsers)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("Printing unmarshlled: ", otherUsers)

	for i, user := range otherUsers {
		fmt.Println("User #", i)
		fmt.Println("\t", user.Age, user.First)
	}
}
