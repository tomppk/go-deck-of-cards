package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Last field contactInfo type contactInfo. Same name and type.
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// bob := person{firstName: "Bob", lastName: "Bobson"}

	// var steve person

	// steve.firstName = "Steve"
	// steve.lastName = "Perry"
	// steve.contact.email = "Steve@journey.com"
	// steve.contact.zipCode = 00123

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Operator &variable Give me the memory address of the value this variable is
	// pointing at
	// So &jim points at the memory address that jim struct is at. It is not a reference to the jim struct but the memory address
	// jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()
}

// *pointer Give me the value this memory address is pointing at
// *pointerToPerson give me the value that is at memory address &jim ie. jim struct
// *person means a type of pointer that points at a person ie. it means we are
// working with a pointer to a person
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// %+v prints all the fields of a struct
func (p person) print() {
	fmt.Printf("%+v", p)

}
