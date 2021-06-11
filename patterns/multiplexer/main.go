package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

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
func sender(from <-chan string, to chan string) {
	for {
		to <- <-from
	}
}
func multiplexer(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go sender(in1, c)
	go sender(in2, c)

	return c
}
func main() {
	ch := multiplexer(
		title("https://www.google.com", "https://www.youtube.com"),
		title("http://www.camboriu.ifc.edu.br", "https://www.toptal.com"),
	)
	fmt.Println("1) ", <-ch, " 2)", <-ch)
	fmt.Println("3) ", <-ch, " 4)", <-ch)
}
