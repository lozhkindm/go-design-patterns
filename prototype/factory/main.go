package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var (
	mainOffice = Employee{
		Name: "",
		Office: Address{
			Suite:  0,
			Street: "123 East Dr",
			City:   "London",
		},
	}
	auxOffice = Employee{
		Name: "",
		Office: Address{
			Suite:  0,
			Street: "66 West Dr",
			City:   "London",
		},
	}
)

type Address struct {
	Suite        int
	Street, City string
}

type Employee struct {
	Name   string
	Office Address
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	employee := proto.DeepCopy()
	employee.Name = name
	employee.Office.Suite = suite
	return employee
}

func (e *Employee) DeepCopy() *Employee {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(e)

	decoder := gob.NewDecoder(&buffer)
	employee := Employee{}
	_ = decoder.Decode(&employee)
	return &employee
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john, jane)
}
