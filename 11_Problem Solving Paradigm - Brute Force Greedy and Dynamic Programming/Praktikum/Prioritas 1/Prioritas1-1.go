package main

import "fmt"

func Biner(n int) string {
	biner := ""

	for n != 1 {
		if n%2 == 0 {
			biner = "1" + biner
			n = n / 2
		} else if n%2 != 0 {
			biner = "0" + biner
			n = n/2 + 1
		}
	}
	return biner
}

func listBiner(n int) []string {
	listbiner := make([]string, 0)
	listbiner = append(listbiner, "0")
	for i := 1; i <= n+1; i++ {
		listbiner = append(listbiner, Biner(i))
	}
	return listbiner
}

func main() {
	fmt.Println(listBiner(9))
}
