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
	oldCanVote := p.CanVote()
	p.age = age
	p.Fire(PropertyChange{
		Name:  "Age",
		Value: p.age,
	})
	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{
			Name:  "CanVote",
			Value: p.CanVote(),
		})
	}
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{
			subs: new(list.List),
		},
		age: age,
	}
}

type ElectoralRoll struct{}

func (e *ElectoralRoll) Notify(data interface{}) {
	if prop, ok := data.(PropertyChange); ok {
		if prop.Name == "CanVote" && prop.Value.(bool) {
			fmt.Println("Congrats, you can vote now")
		}
	}
}

func main() {
	person := NewPerson(0)
	electoralRoll := &ElectoralRoll{}
	person.Subscribe(electoralRoll)

	for i := 10; i < 25; i++ {
		fmt.Println("Setting age to", i)
		person.SetAge(i)
	}
}
