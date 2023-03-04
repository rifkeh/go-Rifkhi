package main

import "fmt"

func fibonnacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fibonnacci(number-2) + fibonnacci(number-1)
	}
}

func main() {
	fmt.Println(fibonnacci(0))  //0
	fmt.Println(fibonnacci(2))  //1
	fmt.Println(fibonnacci(9))  //34
	fmt.Println(fibonnacci(10)) //55
	fmt.Println(fibonnacci(12)) //144
}
