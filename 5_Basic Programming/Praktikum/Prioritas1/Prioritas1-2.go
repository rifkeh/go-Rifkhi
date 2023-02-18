package main

import "fmt"

func main() {
	n := 0
	fmt.Printf("Masukkan sebuah bilangan: ")
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println("Bilangan genap")
	} else {
		fmt.Println("Bilangan ganjil")
	}
}
