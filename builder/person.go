package main

import "fmt"

type Person struct {
	// address
	Street, Postcode, City string

	// job
	Company, Position string
	AnnualIncome      int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{PersonBuilder: *b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{PersonBuilder: *b}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (a *PersonAddressBuilder) Street(street string) *PersonAddressBuilder {
	a.person.Street = street
	return a
}

func (a *PersonAddressBuilder) Postcode(postcode string) *PersonAddressBuilder {
	a.person.Postcode = postcode
	return a
}

func (a *PersonAddressBuilder) City(city string) *PersonAddressBuilder {
	a.person.City = city
	return a
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (j *PersonJobBuilder) Company(company string) *PersonJobBuilder {
	j.person.Company = company
	return j
}

func (j *PersonJobBuilder) Position(position string) *PersonJobBuilder {
	j.person.Position = position
	return j
}

func (j *PersonJobBuilder) AnnualIncome(annualIncome int) *PersonJobBuilder {
	j.person.AnnualIncome = annualIncome
	return j
}

func main() {
	person := NewPersonBuilder().
		Lives().
		Street("Jean Street").
		City("Monaco").
		Postcode("MO74CO").
		Works().
		Company("Ferrari").
		Position("Engineer").
		AnnualIncome(123_000).
		Build()
	fmt.Println(person)
}
