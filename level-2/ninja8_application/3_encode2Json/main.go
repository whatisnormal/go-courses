package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	// Encoded users to a 'writer' impl, in this case os.stdout.
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
