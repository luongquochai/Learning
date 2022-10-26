package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channels with expect variable to share has type string
	channels := make(chan string)

	for _, link := range links {
		/* At here, if using go to create routines (child), it will be a bug
		When we execute command "go run main.go", the func main() is called.
		The main routine is the single routine inside of our program that controls
		when our program exits or quits. So the main routine starts running at some point
		in time as we start to enter the for loop. It starts to create the child go routines
		And at that point, after it finishes the for loop, the main routine is going to exit entirely.
		Even though we still have these other child routines running.
		So the way we're going to do this is by using another construct and go called channels.
		*/
		go checkLink(link, channels)
	}

	// Wait for a value to be sent into the channels. Print out
	// We can receive a lot of messages from channels base on printing out
	// fmt.Println(<-channels)
	// fmt.Println(<-channels) ....
	// But if print the out put > length of links, the system will hang.
	for i := 0; i < len(links); i++ {
		fmt.Println(<-channels)
	}

}

func checkLink(link string, channels chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Send the string "Link is down" into channels
		channels <- "Link is down"
		return
	}

	fmt.Println(link, "is up!")
	channels <- "Link is up"
}
