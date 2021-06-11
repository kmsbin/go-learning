package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{1, "notebook", 2799.00, []string{"sale", "eletronic"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 product
	jsonString := `{"id": 2, "name": "smartphone", "price": 2376.00, "tags": ["2 cameras", "mobile"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
