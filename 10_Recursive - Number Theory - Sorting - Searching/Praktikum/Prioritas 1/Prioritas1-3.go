package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	if number == 1 {
		return 2
	} else {
		count := 1
		biggestPrNum := 0
		for i := 3; count < number; i++ {
			isPrime := true
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isPrime = false
				}
			}
			if isPrime == true {
				count++
				biggestPrNum = i
			}
		}
		return biggestPrNum
	}
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))

}
