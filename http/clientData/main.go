package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"nome_completo"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func main() {
	res, err := http.Get("http://localhost:8000/list-socios")
	// data, _ := json.Unmarshal(res.Body, nill)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var user []user

	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &user)

	fmt.Println(user)
	// fmt.Println(string(body))
}
