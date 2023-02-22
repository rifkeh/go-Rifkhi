package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan angka yang akan dipangkatkan: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scanln(&b)
	fmt.Printf("Hasil pangkatnya adalah %d", pangkat(a, b))

}

func pangkat(a int, b int) int {
	var hasil int
	if b == 0 {
		hasil = 1
	} else if b == 1 {
		hasil = a
	} else if b%2 == 0 {
		hasil = pangkat(a, b/2) * pangkat(a, b/2)
	} else if b%2 != 0 {
		hasil = a * pangkat(a, b/2) * pangkat(a, b/2)
	}
	return hasil
}
