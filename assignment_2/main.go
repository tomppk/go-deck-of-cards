package main

import "fmt"

// For a type to satisfy shape interface, it needs to have method getArea() that returns float64
type shape interface{
	getArea() float64
}

type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}

func main() {
	t := triangle{2,4}
	s := square{5}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}


func printArea(s shape) {
	fmt.Println(s.getArea())
}