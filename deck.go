package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
// The new 'deck' type extends a existing type of []string
// Can think similar to class with methods specific to this class
// as in OOP languages
type deck []string

// A receiver sets up methods on variables we create
// Similar how a class has its own methods in OOP languages
// 'd' is a reference to the variable calling this method eg. 'cards' variable
// Can think of 'd' similar to word 'this' in JS
// By convention this is one or two letters of the type so in this case 'd' 'deck'
// 'deck' is the type we want to attach this method to
// So every variable of type 'deck' gets access to this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}