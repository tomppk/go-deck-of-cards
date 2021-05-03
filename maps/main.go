package main

import "fmt"

func main() {
	// colors := make(map[int]string)

	// var colors map[string]string

	// Map key type string and value type string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// For maps to access and add keys and values use square braces
	// colors[10] = "#fffff"
	// 	delete(colors, 10)

	printMap(colors)
}

// To iterate over map
// for each key, value in map c
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
