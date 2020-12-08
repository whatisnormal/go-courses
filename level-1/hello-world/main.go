package main

// the word package 'main' is reserved for packages that will be executable. With any other word, means that it will not be executable, thus 'resusable'
// first functions also needs to be called 'main'

// give my package main access to all code and functionality in a package called "fmt". Short form of "format". part of the standar library in go.
// If I go to golang.org/pkg, I can see all relevant packages builtin.
import "fmt"

//"func" is short for "functions".
func main() {
	fmt.Println("Hi there!")
}
