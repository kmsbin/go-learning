package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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
func theFaster(curl1, curl2, curl3 string) string {
	c1 := title(curl1)
	c2 := title(curl2)
	c3 := title(curl3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(400 * time.Millisecond):
		return "perderam todos"
	}
}

func main() {
	fmt.Println(theFaster("http://www.camboriu.ifc.edu.br", "https://google.com", "https://youtube.com"))
}
