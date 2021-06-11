package main

import "fmt"

type printable interface {
	toString() string
}

type people struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

func (p people) toString() string {
	return p.name + " " + p.surname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$%.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = people{"kauli", "sabino"}
	fmt.Println(thing.toString())
	print(thing)

	thing = product{"legs", 79.34}
	fmt.Println(thing.toString())
	print(thing)

	p1 := product{"legs", 189.24}
	print(p1)
}
