package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := make([]int, 20, 21)
	arr := [...]int{20, 21}
	fmt.Print(reflect.TypeOf(slice1), reflect.TypeOf(arr))
}
