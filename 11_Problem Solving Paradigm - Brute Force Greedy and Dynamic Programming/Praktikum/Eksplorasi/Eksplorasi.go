package main

import "fmt"

func romawiConvert(n int) string {
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	angka := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	i := 0
	for n > 0 {
		if n >= angka[i] {
			result += romawi[i]
			n -= angka[i]
			i = 0
		} else {
			i++
		}
	}
	return result
}

func main() {
	fmt.Println(romawiConvert(4))
	fmt.Println(romawiConvert(9))
	fmt.Println(romawiConvert(23))
	fmt.Println(romawiConvert(2021))
	fmt.Println(romawiConvert(1646))
}
