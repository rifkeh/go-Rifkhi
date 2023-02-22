package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	var kata string
	boolPalindrome := true
	fmt.Println("Apakah Palindrome?")
	fmt.Printf("Masukkan kata: ")
	scanner.Scan()
	kata = scanner.Text()
	fmt.Println("Captured: " + kata)
	n := len(kata) - 1
	for i := 0; i <= n; i++ {
		if kata[i] != kata[n-i] {
			boolPalindrome = false
		}
	}
	if boolPalindrome == false {
		fmt.Println("Bukan palindrome")
	} else {
		fmt.Println("Palindrome")
	}

}
