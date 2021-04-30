package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	t.Run("assert deck length", func(t *testing.T) {
		d := newDeck()

		got := len(d)
		want := 16

		if got != want {
			t.Errorf("got %v, expected %v", got, want)
		}
	})

	t.Run("first card ace of spades", func(t *testing.T) {
		d := newDeck()

		got := d[0]
		want := "Ace of Spades"

		if got != want {
			t.Errorf("got %v, expected %v", got, want)
		}
	})

		t.Run("last card four of clubs", func(t *testing.T) {
		d := newDeck()

		got := d[len(d) - 1]
		want := "Four of Clubs"

		if got != want {
			t.Errorf("got %v, expected %v", got, want)
		}
	})
}

// Remove "_decktesting" file to cleanup
// Create new deck, save the deck to hard drive.
// Open the deck file and check its length is 16
func TestSaveToDeckAndNewDeckTestFromFile( t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	got := len(newDeckFromFile("_decktesting"))
	want := 16

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	os.Remove("_decktesting")

}