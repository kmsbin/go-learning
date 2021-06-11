package main

import "fmt"

func main() {

	var a = 10
	i := 1

	// Go not has aritmetic to pointers
	var p *int = nil

	p = &i //get i memory adress

	*p++
	i++

	fmt.Println(p, *p, i, a)
}
