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

type Person struct {
	Observable
	Name string
}

func (p *Person) CatchCold() {
	p.Fire(p.Name)
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{
			subs: new(list.List),
		},
		Name: name,
	}
}

type Doctor struct{}

func (d *Doctor) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s", data.(string))
}

func main() {
	person := NewPerson("Billy")
	doctor := &Doctor{}
	person.Subscribe(doctor)
	person.CatchCold()
}
