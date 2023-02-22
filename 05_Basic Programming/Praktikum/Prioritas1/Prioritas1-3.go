package main

import "fmt"

func main() {
	nilai := 0
	grade := ""
	fmt.Printf("Masukkan nilai: ")
	fmt.Scanln(&nilai)
	if nilai <= 34 {
		grade = "E"
	} else if nilai >= 35 && nilai <= 49 {
		grade = "D"
	} else if nilai >= 50 && nilai <= 64 {
		grade = "C"
	} else if nilai >= 65 && nilai <= 79 {
		grade = "B"
	} else if nilai >= 80 && nilai <= 100 {
		grade = "A"
	} else {
		fmt.Println("Masukkan nilai hanya dengan bilangan positif")
	}
	fmt.Println("Grade anda adalah " + grade)
}
