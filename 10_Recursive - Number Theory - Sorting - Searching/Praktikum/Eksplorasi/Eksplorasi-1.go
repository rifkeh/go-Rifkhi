package main

import "fmt"

func maxSequence(arr []int) int {
	maxTemp := arr[0]
	maxFinal := arr[0]

	for i := 1; i < len(arr); i++ {
		// Mencari nilai maximal antara maxTemp ditambah next array atau hanya next array
		maxTemp = max(arr[i], maxTemp+arr[i])
		// Mencari nilai maximal antara maxFinal sebelumnya dengan maxTemp baru
		maxFinal = max(maxFinal, maxTemp)
	}

	return maxFinal
}

// Fungsi mencari nilai maksimum antara dua integer
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(maxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(maxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
	fmt.Println(maxSequence([]int{5, 5, 5, 5, 5, 5, 5}))           // 35
}
