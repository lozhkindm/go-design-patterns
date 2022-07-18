package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

type tiredPerson struct {
	name string
	age  int
}

func (t *tiredPerson) SayHello() {
	fmt.Println("Sorry, I'm too tired")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{
			name: name,
			age:  age,
		}
	}
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	person := NewPerson("John", 22)
	person.SayHello()

	tired := NewPerson("Michael", 121)
	tired.SayHello()
}
