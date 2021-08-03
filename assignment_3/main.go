package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args is a variable that represents the command line arguments passed into function of type []string
	// [0] is the path to temporary file that gets created. [1] is the text added after go run main.go "textHere"

	textFile := os.Args[1]

	// os.Open opens textFile at path from os.Args[1] and reads content and writes it into file variable
	// io.Copy takes content from file and reads it and writes it into os.Stdout ie. terminal
	file, err := os.Open(textFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}