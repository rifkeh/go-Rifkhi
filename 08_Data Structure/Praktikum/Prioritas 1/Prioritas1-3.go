package main

import (
	"fmt"
	"strconv"
)

func main() {
	var angka string
	fmt.Print("Masukkan angka yang akan di cek muncul sekali nya: ")
	fmt.Scanln(&angka)
	fmt.Println(MunculSekali(angka))
}

func MunculSekali(angka string) []int {
	var isSame bool
	var hasil []int
	for i, temp := range angka {
		isSame = false
		for j, temp1 := range angka {
			if temp == temp1 && i != j {
				isSame = true
			}
		}
		if isSame == false {
			intVar, _ := strconv.Atoi(string(temp))
			hasil = append(hasil, intVar)
		}
	}
	return hasil
}
