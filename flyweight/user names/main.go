package main

import (
	"fmt"
	"strings"
)

var allNames []string

type User struct {
	names []uint8
}

func (u *User) GetFullName() string {
	var parts []string
	for _, name := range u.names {
		parts = append(parts, allNames[name])
	}
	return strings.Join(parts, " ")
}

func NewUser(fullName string) *User {
	getOrAdd := func(name string) uint8 {
		for i := range allNames {
			if allNames[i] == name {
				return uint8(i)
			}
		}
		allNames = append(allNames, name)
		return uint8(len(allNames) - 1)
	}

	user := User{}
	parts := strings.Split(fullName, " ")
	for _, part := range parts {
		user.names = append(user.names, getOrAdd(part))
	}
	return &user
}

func main() {
	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	smith := NewUser("Jane Smith")
	fmt.Println(john.GetFullName(), jane.GetFullName(), smith.GetFullName())
	fmt.Println(john.names, jane.names, smith.names)

	totalMem := 0
	for _, name := range allNames {
		totalMem += len([]byte(name))
	}
	totalMem += len(john.names)
	totalMem += len(jane.names)
	totalMem += len(smith.names)
	fmt.Println(totalMem)
}
