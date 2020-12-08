package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, error := os.Open(fileName)

	if error != nil {
		fmt.Println("Error opening file name: ", fileName, error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
