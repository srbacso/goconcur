package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Printf("Type something and press enter (enter Q to quit)")
	for {
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if "q" == strings.ToLower(userInput) {
			break
		}
		ping <- userInput
		response := <-pong
		fmt.Println("Response: ", response)
	}
	fmt.Println("All done close channels")
	close(ping)
	close(pong)
}
