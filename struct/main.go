package main

import "log"

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
	log.Println(person)

	person.firstName = "Jane"
	person.lastName = "Doe"
	person.contact = contactInfo{"janedoe@example.com", 67890}
	log.Println(person)
}

