package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(person *Person)

type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Name(name string) *PersonBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.name = name
	})
	return b
}

func (b *PersonBuilder) Position(position string) *PersonBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.position = position
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	person := Person{}
	for _, action := range b.actions {
		action(&person)
	}
	return &person
}

func main() {
	builder := PersonBuilder{}
	person := builder.
		Position("Developer").
		Name("John").
		Build()
	fmt.Println(person)
}
