package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional

func NewEmployeeFactoryFunctional(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

// structural

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (e *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		Name:         name,
		Position:     e.Position,
		AnnualIncome: e.AnnualIncome,
	}
}

func NewEmployeeFactoryStructural(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func main() {
	developerFactory := NewEmployeeFactoryFunctional("developer", 60000)
	managerFactory := NewEmployeeFactoryFunctional("manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("James")

	fmt.Println(developer, manager)

	ceoFactory := NewEmployeeFactoryStructural("ceo", 100000)
	ceoFactory.AnnualIncome = 110000
	ceo := ceoFactory.Create("Jacob")
	fmt.Println(ceo)
}
