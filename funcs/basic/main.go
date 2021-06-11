package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s", p1, p2)
}
func f3() string {
	return "É nois"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf(" F4: %s, %s", p1, p2)
}

func f5(p1 string) (string, string) {
	fmt.Println("i can return multiples values ")
}

func main() {
	fmt.Println("jed")
}
