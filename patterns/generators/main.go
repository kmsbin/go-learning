package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - read only channel

func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(curl string) {
			resp, _ := http.Get(curl)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			title := r.FindStringSubmatch(string(html))

			if len(title) == 0 {
				fmt.Println(title)
				c <- "erro no link " + curl
				return
			}
			c <- title[1]
		}(url)
	}
	return c
}
func main() {
	t1 := title("https://www.google.com", "http://www.camboriu.ifc.edu.br")
	fmt.Println("primeiros", <-t1, "|", <-t1)
	// fmt.Println("segundos", <-t1, "|", <-t2)
}
