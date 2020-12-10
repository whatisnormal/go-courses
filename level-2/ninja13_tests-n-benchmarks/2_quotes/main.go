package main

import (
	"example.com/username/repo/ninja13_tests-n-benchmarks/2_quotes/quote"
	"example.com/username/repo/ninja13_tests-n-benchmarks/2_quotes/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
