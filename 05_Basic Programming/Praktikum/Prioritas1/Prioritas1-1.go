package main

import "fmt"

func main() {
	var sisiA, sisiB, tinggi, luas float32 = 0, 0, 0, 0
	fmt.Printf("Masukkan panjang salah satu sisi yang berhadapan: ")
	fmt.Scanln(&sisiA)
	fmt.Printf("Masukkan panjang satu sisinya lagi yang berhadapan: ")
	fmt.Scanln(&sisiB)
	fmt.Printf("Masukkan tinggi trapesium: ")
	fmt.Scanln(&tinggi)
	luas = (sisiA + sisiB) * (tinggi / 2)
	fmt.Printf("Luas trapesium adalah ")
	fmt.Println(luas)
}
