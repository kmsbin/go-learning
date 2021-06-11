package main

import "fmt"

type grades float64

func (g grades) in(begin, end float64) bool {
	return float64(g) >= begin && float64(g) <= end
}

func grade2concept(g grades) string {

	switch {
	case g.in(9.0, 10.0):
		return "A"
	case g.in(7.0, 8.99):
		return "B"
	case g.in(5.0, 7.99):
		return "C"
	case g.in(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	var grade grades = 7.0

	fmt.Println(grade2concept(grade))
}
