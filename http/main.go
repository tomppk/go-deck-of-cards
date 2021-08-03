package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{} 

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// Same thing as below but with only one line
	// First argument is the output channel where to copy the information, second argument is the input channel where the information comes from
	// First argument something that implements Writer interface, second argument something that implements Reader interface
	io.Copy(lw, resp.Body)

	// make() is a built-in function that takes a type of a slice and second argument is the number of elements or empty spaces that the slice will be initialized with
	// bs := make([]byte, 99999)

	// Take the Body of the response and pass it to Read() function that will read the data and pass it into the []byte that was just created
	// resp.Body.Read(bs)

	// Turn bs into string before printing it out
	// fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	// To implement the Writer interface Write() must return int number of bytes read and an error
	return len(bs), nil
}