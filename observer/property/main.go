package main

import (
	"container/list"
	"fmt"
)

type Observer interface {
	Notify(data interface{})
}

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(observer Observer) {
	o.subs.PushFront(observer)
}

func (o *Observable) Unsubscribe(observer Observer) {
	for i := o.subs.Front(); i != nil; i = i.Next() {
		if i.Value.(Observer) == observer {
			o.subs.Remove(i)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for i := o.subs.Front(); i != nil; i = i.Next() {
		i.Value.(Observer).Notify(data)
	}
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	age int
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if p.age == age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{
		Name:  "Age",
		Value: p.age,
	})
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{
			subs: new(list.List),
		},
		age: age,
	}
}

type TrafficManagement struct {
	person Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if prop, ok := data.(PropertyChange); ok {
		if prop.Value.(int) >= 18 {
			fmt.Println("Congrats, you can drive now")
			t.person.Unsubscribe(t)
		}
	}
}

func main() {
	person := NewPerson(15)
	trafficManagement := &TrafficManagement{person: person.Observable}
	person.Subscribe(trafficManagement)

	for i := 15; i <= 20; i++ {
		fmt.Println("Setting the age to", i)
		person.SetAge(i)
	}
}
