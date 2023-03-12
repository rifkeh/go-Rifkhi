package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i * 3
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}

}
