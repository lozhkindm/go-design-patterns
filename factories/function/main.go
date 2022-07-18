package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}
}

func main() {
	person := NewPerson("John", 22)
	fmt.Println(person)
}
