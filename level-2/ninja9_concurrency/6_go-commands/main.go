package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)

	/*
	   If you want to install your binary, run
	   1) on the project folder -> 'go build' and 'go install'
	   2) the previous command will install the binary on goworkspace/bin

	*/
}
