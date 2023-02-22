package main

import (
	"fmt"
	"math"
)

func main() {
	isTrue := true
	var n float64
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&n)
	for i := 2; i <= int(math.Sqrt(n)); i++ {
		if int(n)%i == 0 {
			isTrue = false
		}
	}
	if isTrue == false {
		fmt.Println("Bilangan tidak prima")
	} else {
		fmt.Println("Bilangan Prima")
	}
}
