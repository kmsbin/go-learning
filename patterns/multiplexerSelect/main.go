package main

import (
	"fmt"
	"time"
)

func talk(people string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", people, i)
		}
	}()
	return c
}
func multiplexer(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-in1:
				c <- s
			case s := <-in2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	ch := multiplexer(talk("kauli"), talk("sajdf"))
	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
}
