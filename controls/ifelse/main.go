package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printResult(resul float64) string {

	if resul > 7 {
		return "congratulations you had approved " + fmt.Sprint(resul)
	} else {
		return "sorry try again next year " + fmt.Sprint(resul)
	}
}

func randomNumb() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Println(randomNumb())
		time.Sleep(time.Millisecond)
	}

}
