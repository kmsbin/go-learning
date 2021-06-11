package main

import "fmt"

func main() {
	numbs := [...]int{1, 2, 3, 4, 5}

	for i, numb := range numbs {
		fmt.Printf("%d) %d\n", i+1, numb)
	}
}
