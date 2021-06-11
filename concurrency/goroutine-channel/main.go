package main

import (
	"fmt"
	"time"
)

// Channel syntax <- output, -> input

func numbTimer(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second * 3)
	c <- 4 * base
}
func main() {
	c := make(chan int)
	go numbTimer(4, c)

	a, b := <-c, <-c
	fmt.Println(a, b)
	fmt.Println(<-c)

}
