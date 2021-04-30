package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Takes a deck and returns a string representation of it
func (d deck) toString() string {
	// Type conversion, turn deck type into []string
	// From "strings" package join []strings together separated by comma
	return strings.Join([]string(d), ",")
}

// Type conversion, turn deck into one long string, and turn that string into
// slice of bytes so WriteFile can write it to hard drive
// Last argument default permission 0666 meaning anyone can read/write the file
// WriteFile() returns an error if something went wrong
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read file and return deck
// ReadFile() returns []byte and error. If error log it and quit program
func newDeckFromFile(filename string) deck {
	bs, err :=	ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Type conversion, convert byte slice into a string of cards separated by
	// commas ie. Ace of Spades,Four of Hearts,...
	// Split that string at commas to get s slice of strings []string
	// Convert that s into type deck

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// index of card
// Generate random number between 0 and len of slice -1
// Truly random number using Source and rand.New()
// source = generate new Source of randomness. use time.Now().UnixNano() to
// get a random int64 value each iteration.
// as Source is a random seed every time then r = new rand Random type with
// random seed value each time to generate truly random numbers

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// iterate over all elements in deck
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Assign value that is at index d[newPosition] to d[i]
		// and value that is at d[i] to newPosition
		// so swap positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}