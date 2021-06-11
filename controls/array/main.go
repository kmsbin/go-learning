package main

import "fmt"

func main() {
	var note [3]float64

	note[0], note[1], note[2] = 4.3, 4.5, 10.0

	fmt.Println(note)

	total := 0.0
	for i := 0; i < len(note); i++ {
		total += note[i]
	}
	fmt.Println(total / 3)
}
