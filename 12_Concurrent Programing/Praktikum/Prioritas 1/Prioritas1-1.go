package main

import (
	"fmt"
	"time"
)

func multiple(x int) {
	for i := 1; i >= 0; i++ {
		fmt.Println(x * i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go multiple(3)
	fmt.Println("test")
	time.Sleep(18 * time.Second)
}
