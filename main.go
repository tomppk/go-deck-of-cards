package main

import "fmt"


func main() {
	// The Go compiler will infer types for you with :=
	// := only used when variable is declared the first time, the very first time
	// variable is initialized. When re-declaring variable value use just =

	// Defining new variable and manually assigning type
	// var card string = "Ace of Spades"

	card := newCard()
	fmt.Println(card)

	card = "Five of Diamonds"
	fmt.Println(card)

}

// The function will return type 'string'
func newCard() string {
	return "Queen of Hearts"
}
