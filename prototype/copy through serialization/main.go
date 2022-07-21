package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(p)

	fmt.Println(string(buffer.Bytes()))

	decoder := gob.NewDecoder(&buffer)
	person := Person{}
	_ = decoder.Decode(&person)
	return &person
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
