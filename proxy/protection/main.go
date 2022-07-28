package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    *Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		car:    &Car{},
		driver: driver,
	}
}

func main() {
	car := NewCarProxy(&Driver{Age: 12})
	car.Drive()

	car2 := NewCarProxy(&Driver{Age: 22})
	car2.Drive()
}
