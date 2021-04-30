package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
// The new 'deck' type extends a existing type of []string
// Can think similar to class with methods specific to this class
// as in OOP languages
type deck []string

// Function that returns type deck
// Create new variable cards of type deck
func newDeck() deck {
	cards := deck{}

	// Slice of strings. Not a type deck.
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Loop over list of suits and for each suit loop over list of values
	// and for each value append value + suit together as "Ace of Spades"
	// Can replace indexes i and j with underscore _ to tell go that these
	// variables will be unused
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// A receiver sets up methods on variables we create
// Similar how a class has its own methods in OOP languages
// 'd' is a reference to the variable calling this method eg. 'cards' variable
// Can think of 'd' similar to word 'this' in JS
// By convention this is one or two letters of the type so in this case 'd' 'deck'
// 'deck' is the type we want to attach this method to
// So every variable of type 'deck' gets access to this method
// range d ie. range 'cards'. Range is a keyword we use whenever we want to
// iterate over every item in a slice
// i, card =: range d assigns variables index and card to every element in a slice
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return multiple values. Two values type deck.
// First return value from index 0 to handSize and second from handSize to last i
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
