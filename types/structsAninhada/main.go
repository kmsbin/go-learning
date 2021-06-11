package main

import "fmt"

type item struct {
	product int
	qtd     int
	price   float64
}
type order struct {
	userID int
	items  []item
}

func (o order) priceOrder() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.qtd)
	}
	return total
}

func main() {
	order := order{
		userID: 1,
		items: []item{
			item{1, 2, 25.53},
			item{2, 53, 23.23},
			item{3, 23, 75.84},
			item{4, 6, 36.5},
			item{5, 12, 345.6},
		},
	}

	fmt.Println("total price is", order.priceOrder())
}
