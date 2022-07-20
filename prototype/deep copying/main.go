package main

import "fmt"

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{
		Name: "John",
		Address: &Address{
			Street:  "123 London Rd",
			City:    "London",
			Country: "UK",
		},
	}

	jane := john
	jane.Name = "Jane"                   // ok
	jane.Address.Street = "321 Baker St" // problem

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	// deep copying
	margo := jane
	margo.Name = "Margo"
	margo.Address = &Address{
		Street:  jane.Address.Street,
		City:    jane.Address.City,
		Country: jane.Address.Country,
	}
	margo.Address.Street = "123 London Rd"

	fmt.Println(jane, jane.Address)
	fmt.Println(margo, margo.Address)
}
