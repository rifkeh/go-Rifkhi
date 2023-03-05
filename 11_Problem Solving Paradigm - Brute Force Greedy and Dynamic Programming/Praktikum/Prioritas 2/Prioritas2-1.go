package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	temp := make([]int, len(jumps))
	temp[0] = 0
	temp[1] = int(math.Abs(float64(jumps[1]) - float64(jumps[0])))
	for i := 2; i <= len(jumps)-1; i++ {
		oneStep := temp[i-1] + int(math.Abs(float64(jumps[i])-float64(jumps[i-1])))
		twoStep := temp[i-2] + int(math.Abs(float64(jumps[i])-float64(jumps[i-2])))
		temp[i] = int(math.Min(float64(oneStep), float64(twoStep)))
	}
	return temp[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
