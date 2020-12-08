package hello

import "rsc.io/quote"


// Add this as a module with:
// go mod init "example.com/username/repo"
func Hello() string {
	
	return quote.Hello()
}
