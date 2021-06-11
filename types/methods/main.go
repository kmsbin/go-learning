package main

import (
	"fmt"
	"strings"
)

type people struct {
	name    string
	surname string
}

func (p people) getCompleteName() string {
	return p.name + " " + p.surname
}

func (p *people) setCompleteName(completedName string) {
	slice := strings.Split(completedName, " ")
	p.name = slice[0]
	p.surname = slice[1]
}

func main() {
	people := people{"kauli", "sabino"}
	fmt.Println(people.getCompleteName())

	people.setCompleteName("Tyrion Lennister")
	fmt.Println(people.getCompleteName())
}
