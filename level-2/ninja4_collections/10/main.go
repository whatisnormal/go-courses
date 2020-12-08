package main

import "fmt"

func main() {

	val1 := []string{"James", "Bond", "Shaken", "Not stirred"}
	val2 := []string{"Miss", "Moneypenny", "Hellooo james."}

	m := map[string][]string{
		`bond_james`:      val1,
		`moneypenny_miss`: val2,
	}

	m["jose_joana"] = []string{"jose", "joana"}

	delete(m, "bond_james")

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

	}

}
