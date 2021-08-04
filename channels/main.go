package main

import (
	"fmt"
	"net/http"
	"time"
)

// Program to check website status. Make http request and print if site is up or down.
func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel of type string to communicate string information between go routines.
	c := make(chan string)

	// go keyword creates a new go routine that runs the checkLink function. 
	for _, link := range links {
		go checkLink(link, c)
	}

	// Wait for a value to be sent into a channel. When we get a value, print it out immediately. Value coming through channel is a blocking call. We are putting this for loop together just to wait for every Go routine to emit a message into a channel and then print it out.
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<- c)
	// }

	// The for loop means, watch the channel c and whenever it emits a value, assign it to variable l short for link. Once a value comes in, the body of the for loop is immediately executed so start a new go routine and call checkLink function. First argument is the link url that is sent into the channel, second is the channel.
	for l := range c {
		// Function literal ie. anonymous function. Add extra set of parentheses at the end to actually invoke the function, add link as an argument to function literal so it gets access to copy of l from the main go routine ie. the for loop. Function l refers to copy, not the same address in memory as main routine or for loop l
		go func(link string) {
			// Pause the current go routine for five seconds
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// Make HTTP request to link. We get back two values: struct response object and error if one occurred. We do not care about response, we just want to check if the website is down or not ie if there is an error. Add second argument of type channel that communicates string data between go routines. So between func main parent go routine that gets created automatically and the child go routines that are created inside for loop
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Send string into a channel ie. the url of link
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}