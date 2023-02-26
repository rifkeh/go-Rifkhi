package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var angka string
	// fmt.Print("Masukkan angka yang akan di cek muncul sekali nya: ")
	// fmt.Scanln(&angka)
	// fmt.Println(MunculSekali(angka))
	fmt.Println(MunculSekali("1234123"))
	fmt.Println(MunculSekali("76523752"))
	fmt.Println(MunculSekali("12345"))
	fmt.Println(MunculSekali("1122334455"))
	fmt.Println(MunculSekali("0872504"))
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
