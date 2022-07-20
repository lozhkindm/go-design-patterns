package main

import "fmt"

type Address struct {
	Street, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		Street:  a.Street,
		City:    a.City,
		Country: a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	pp := *p
	pp.Address = p.Address.DeepCopy()
	copy(pp.Friends, p.Friends)
	return &pp
}

func main() {
	john := Person{
		Name: "John",
		Address: &Address{
			Street:  "123 London Rd",
			City:    "London",
			Country: "UK",
		},
		Friends: []string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.Street = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
