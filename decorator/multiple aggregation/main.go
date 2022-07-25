package main

import "fmt"

type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling")
	}
}

type Dragon struct {
	Bird
	Lizard
}

func (d *Dragon) GetAge() int {
	return d.Bird.Age
}

func (d *Dragon) SetAge(age int) {
	d.Bird.Age = age
	d.Lizard.Age = age
}

func main() {
	dragon := Dragon{}
	dragon.SetAge(5)
	dragon.Fly()
	dragon.Crawl()
}
