package main

import "fmt"

// For type to satisfy bot interface, it needs to have method getGreeting() that
// returns a string
// ie. it needs to have a receiver function
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// Function to be used with a type that satisfies the bot interface
// Both english and spanish bots satisfy the interface as they have getGreeting()
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
