package main

import "fmt"

type (
	Person struct {
		firstName string
		lastName  string
		ContactInfo
	}

	ContactInfo struct {
		email   string
		zipCode int
	}
)

func main() {

	person := Person{
		firstName: "John",
		lastName:  "Doe",
		ContactInfo: ContactInfo{
			email:   "johnDoe@gmail.com",
			zipCode: 12345,
		},
	}

	person.updateName("Jane")

	person.print()
}

func (p *Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(firstName string) {
	p.firstName = firstName
}
