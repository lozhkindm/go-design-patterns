package main

import "fmt"

const (
	Developer = iota
	Manager
)

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{
			Position:     "developer",
			AnnualIncome: 60000,
		}
	case Manager:
		return &Employee{
			Position:     "manager",
			AnnualIncome: 80000,
		}
	default:
		return &Employee{}
	}
}

func main() {
	manager := NewEmployee(Manager)
	manager.Name = "Jason"
	fmt.Println(manager)
}
