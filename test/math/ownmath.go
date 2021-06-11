package ownmath

import (
	"fmt"
	"strconv"
)

//Media é responsável por calcular a média aritmética
func Media(numbs ...float64) float64 {
	var total float64 = 0.0
	for _, numb := range numbs {
		total += numb
	}

	media := total / float64(len(numbs))

	aroundedMedia, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return aroundedMedia
}
