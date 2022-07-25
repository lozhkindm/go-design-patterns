package main

import "fmt"

type Aged interface {
	GetAge() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) GetAge() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) GetAge() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) GetAge() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.age = age
	d.lizard.age = age
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{
		bird:   Bird{},
		lizard: Lizard{},
	}
}

func main() {
	dragon := NewDragon()
	dragon.SetAge(5)
	dragon.Fly()
	dragon.Crawl()
}
