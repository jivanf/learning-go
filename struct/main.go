package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	person := person{
		firstName: "John",
		lastName: "Doe",
		contact: contactInfo {
			"johndoe@example.com", 12345,
			},
	}
	person.updateFirstName("Smith")
	person.print()

	person.firstName = "Jane"
	person.lastName = "Doe"
	person.contact = contactInfo{"janedoe@example.com", 67890}
	person.print()
}
func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}
