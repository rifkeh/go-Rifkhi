package main

import "fmt"

func main() {
	var idx, target int
	fmt.Print("Masukkan berapa banyak bilangan: ")
	fmt.Scanln(&idx)
	arr := make([]int, idx)
	for i := 0; i <= idx-1; i++ {
		fmt.Printf("Masukkan bilangan ke-%v: ", i+1)
		fmt.Scanln(&arr[i])
	}
	fmt.Printf("Masukkan target penjumlahan: ")
	fmt.Scanln(&target)
	fmt.Print("Anggota array yang memenuhi penjumlahan menjadi target adalah: ")
	fmt.Println(PairSum(arr, target))
}

func PairSum(arr []int, target int) [][]int {
	var hasil [][]int
	var temp []int
	for i, slice1 := range arr {
		for j, slice2 := range arr {
			if slice1+slice2 == target && i < j {
				temp = nil
				temp = append(temp, i, j)
				hasil = append(hasil, temp)
			}
		}
	}

	return hasil
}
