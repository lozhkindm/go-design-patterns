package main

import "fmt"

type Person struct {
	Name string
	Room *ChatRoom
	logs []string
}

func (p *Person) Receive(sender, message string) {
	msg := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s\n", p.Name, msg)
	p.logs = append(p.logs, msg)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(to, message string) {
	p.Room.Message(p.Name, to, message)
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(sender, message string) {
	for _, p := range c.people {
		if p.Name != sender {
			p.Receive(sender, message)
		}
	}
}

func (c *ChatRoom) Message(sender, receiver, message string) {
	for _, p := range c.people {
		if p.Name == receiver {
			p.Receive(sender, message)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	message := fmt.Sprintf("%s joins the chat", p.Name)
	c.Broadcast("Room", message)
	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hello chat")
	jane.Say("hi")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("hey")

	jane.PrivateMessage("Simon", "hi, Simon")
}
