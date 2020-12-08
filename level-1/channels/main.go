package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Making a channel of type 'string'
	// A channel is a 2-way sending device.
	// channel <- 5 Sends the value 5 into this channel
	// myNumber <- channel Waits for a value to be sent into the channel, when we got one, assign the value to 'myNumber'
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Receive from the channel
	// it is a blocking instruction.
	//fmt.Println(<-c)

	/*
		for i := 0; i < len(links); i++ {
			fmt.Println(<-c)
		}
	*/

	// Infinite loop
	/*
		for {
			go checkLink(<-c, c)
		}
	*/

	// Correct syntax for reading from channels.
	// l, short for link.
	for l := range c {
		// Pause for 5 seconds
		//time.Sleep(5 * time.Second)
		//go checkLink(l, c) //replace with function literal ( anonymous function )

		//Function literal / anonymous function / lambda
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // variable passed from the 'for'
	}
}

func checkLink(link string, c chan string) {
	// As we do not need the response, we can use '_'
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link

}
