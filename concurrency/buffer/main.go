package main

import (
	"fmt"
	"time"
)

func rotine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println(".EXE")
	ch <- 5
	ch <- 6
}
func main() {
	ch := make(chan int, 6)
	go rotine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

}
