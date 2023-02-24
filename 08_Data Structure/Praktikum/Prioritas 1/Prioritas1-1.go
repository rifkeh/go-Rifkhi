package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	fmt.Print("Masukkan banyaknya anggota slice pertama: ")
	fmt.Scanln(&a)
	slice1 := make([]string, a, a)
	for i := 0; i <= a-1; i++ {
		fmt.Printf("Masukkan anggota ke-%v slice pertama: ", i+1)
		fmt.Scanln(&slice1[i])
	}
	fmt.Print("Masukkan banyaknya anggota slice kedua: ")
	fmt.Scanln(&b)
	slice2 := make([]string, b, b)
	for i := 0; i <= b-1; i++ {
		fmt.Printf("Masukkan anggota ke-%v slice pertama: ", i+1)
		fmt.Scanln(&slice2[i])
	}
	fmt.Println(strings.Join(arrayMerge(slice1, slice2), ", "))
}

func arrayMerge(a, b []string) []string {
	var c []string
	var isSame bool
	for _, slice := range a {
		c = append(c, slice)
	}
	for _, slice2 := range b {
		isSame = false
		for _, slice := range a {
			if slice2 == slice {
				isSame = true
			}
		}
		if isSame == false {
			c = append(c, slice2)
		}
	}
	return c
}
