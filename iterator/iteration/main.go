package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	names := make(chan string)
	go func() {
		defer close(names)
		names <- p.FirstName
		if len(p.MiddleName) > 0 {
			names <- p.MiddleName
		}
		names <- p.LastName
	}()
	return names
}

type PersonNameIterator struct {
	person  *Person
	current int
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	default:
		panic("invalid current index")
	}
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person: person, current: -1}
}

func main() {
	p := Person{
		FirstName:  "Alexander",
		MiddleName: "Graham",
		LastName:   "Bell",
	}
	for _, name := range p.Names() {
		fmt.Println(name)
	}

	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}

	for iter := NewPersonNameIterator(&p); iter.MoveNext(); {
		fmt.Println(iter.Value())
	}
}
