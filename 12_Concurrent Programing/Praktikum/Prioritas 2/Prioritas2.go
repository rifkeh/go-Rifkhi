package main

import (
	"fmt"
	"sync"
)

type word struct {
	value []int
	mu    sync.Mutex
}

func (w *word) searchChar(strings string, char byte, wg *sync.WaitGroup) {
	// w.mu.Lock()
	for i := range strings {
		if char == strings[i] {
			w.value[char-97] += 1
		}
	}
	// w.mu.Unlock()
	wg.Done()
}

func main() {
	var kata word
	const alphabet, kalimat = "abcdefghijklmnopqrstuvwxyz", "rifkhiakbar"
	kata.value = make([]int, len(alphabet))
	var wg sync.WaitGroup
	for i := range alphabet {
		wg.Add(1)
		go kata.searchChar(kalimat, alphabet[i], &wg)
	}
	wg.Wait()
	for i := range alphabet {
		fmt.Print(string(rune(i+97)) + ": ")
		fmt.Println(kata.value[i])
	}
}
