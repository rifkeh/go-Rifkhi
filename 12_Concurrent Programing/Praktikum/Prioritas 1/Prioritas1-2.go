package main

import "fmt"

func main() {
	ch := make(chan int)

	for j := 1; j <= 10; j++ {
		go func() {
			i := <-ch
			fmt.Println(3 * i)
		}()
		ch <- j
	}
}
