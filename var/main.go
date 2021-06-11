package main

import (
	f "fmt"
	"math"
	"reflect"
)

func main() {
	float64 PI = 3.1415

	// x := 3.14153463
	// var raio = 3.2
	// area := PI * m.Pow(raio, 2)
	f.Printf("safas %f", PI)
	f.Println("\na area é " + f.Sprint(PI))

	var mI = math.MaxFloat64

	f.Println("O tipo é", reflect.TypeOf(mI))

}
