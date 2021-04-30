package main

func main() {
	// The Go compiler will infer types for you with :=
	// := only used when variable is declared the first time, the very first time
	// variable is initialized. When re-declaring variable value use just =

	// Defining new variable and manually assigning type
	// var card string = "Ace of Spades"

	// card := newCard()
	// fmt.Println(card)

	// card = "Five of Diamonds"
	// fmt.Println(card)

	// Declare a slice and type it contains. Inside curly braces {} add the
	// content the slice contains.
	// append() does not modify existing slice but creates a new slice that
	// we assign back to variable cards
	// cards := newDeck()
	// cards = append(cards, "Six of Spades")

	// Iterate over a closed set or a slice of records
	// for index, current card we are iterating over.
	// range is keyword used to iterate over every record in a slice
	// for loops use := syntax because for every loop the index and card
	// need to be re-initialized as they are thrown away at the end of loop
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()

	// Deal returns two values of type deck. Assign them to two variables
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards := newDeckFromFile("my_cards")
	cards.print()		
}

// The function will return type 'string'
func newCard() string {
	return "Queen of Hearts"
}
