package main

import "fmt"

func fibo(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fibo(number-2) + fibo(number-1)
	}
}

func main() {
	fmt.Println(fibo(0))  //0
	fmt.Println(fibo(2))  //1
	fmt.Println(fibo(9))  //34
	fmt.Println(fibo(10)) //55
	fmt.Println(fibo(12)) //144
}
