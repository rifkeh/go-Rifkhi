package main

import (
	"fmt"
	"math"
)

func main() {
	var idx int
	fmt.Print("Masukkan panjang matrix: ")
	fmt.Scanln(&idx)
	// temp1 := make([]int, idx)
	matrix := make([][]int, idx)
	for i := 0; i <= idx-1; i++ {
		matrix[i] = make([]int, idx)
		for j := 0; j <= idx-1; j++ {
			fmt.Printf("Masukkan nilai untuk baris ke-%v dan kolom ke-%v: ", i, j)
			fmt.Scanln(&matrix[i][j])
		}
	}

	for i := 0; i <= idx-1; i++ {
		fmt.Println(matrix[i])
	}
	fmt.Print("Selisih antar diagonal adalah: ")
	fmt.Println(SelisihAbsolut(matrix))
}

func SelisihAbsolut(matrix [][]int) int {
	var selisih1, selisih2 int
	for i := 0; i <= len(matrix)-1; i++ {
		selisih1 += matrix[i][i]
		selisih2 += matrix[i][len(matrix)-i-1]
	}
	return int(math.Abs(float64(selisih1 - selisih2)))
}
