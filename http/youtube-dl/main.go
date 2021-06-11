package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.youtube.com/watch?v=sONL6QUMR9E&ab_channel=Megadeth-Topic")
	if err != nil {
		log.Fatal("FUDEU ")

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	_ = ioutil.WriteFile("file.json", body, 0644)

	fmt.Println(string(body))

}
