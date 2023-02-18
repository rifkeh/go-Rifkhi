package main

import "fmt"

func main() {
	n := 0
	fmt.Printf("Masukkan sebuah angka: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}
