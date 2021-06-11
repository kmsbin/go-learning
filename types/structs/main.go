package main

import (
	"fmt"
	"reflect"
)

type product struct {
	name      string
	price     float64
	promotion float64
}

// recipe method
func (p product) priceWithPromotion() float64 {
	return p.price * (1 - p.promotion)
}
func main() {
	var p1 product
	p1 = product{
		name:      "bola",
		price:     25.0,
		promotion: 0.1,
	}
	fmt.Println(reflect.TypeOf(p1))
}
