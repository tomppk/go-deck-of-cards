package main

import "testing"

func TestList(t *testing.T) {
	list := newList()

	got := len(list)
	want := 11

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
