package main

import (
	"fmt"
)

func rotine(c chan int) {
	// time.Sleep(time.Second)
	c <- 1
	fmt.Println("depois de ser lido")
}

func main() {
	var c chan int
	go rotine(c)

	fmt.Println(<-c)
	fmt.Println("foi lido")
	// fmt.Println(<-c)

}
