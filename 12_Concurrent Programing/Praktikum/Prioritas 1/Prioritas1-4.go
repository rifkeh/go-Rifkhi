package main

import (
	"fmt"
	"sync"
)

type number struct {
	value int
	mu    sync.Mutex
}

func (n *number) factorialNumber(num int, w *sync.WaitGroup) {
	n.mu.Lock()
	n.value *= num
	w.Done()
	n.mu.Unlock()
}

func (n *number) getNumber() int {
	return n.value
}

func main() {
	var angka number
	angka.value = 1
	var wg sync.WaitGroup
	for i := 2; i <= 5; i++ {
		wg.Add(1)
		go angka.factorialNumber(i, &wg)
	}

	wg.Wait()
	fmt.Println(angka.getNumber())

}
