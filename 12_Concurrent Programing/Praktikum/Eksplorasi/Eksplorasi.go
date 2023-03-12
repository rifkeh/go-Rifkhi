package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type Products struct {
	Product []Product
	mu      sync.Mutex
}

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

// func (p *Products) readData(ch chan interface{}, wg *sync.WaitGroup) {
// 	for i := 0; i <= len(p.Product); i++ {
// 		p.mu.Lock()
// 		ch <- p.Product[i].C
// 		p.mu.Unlock()
// 	}
// }

func main() {
	const url = "https://fakestoreapi.com/products"
	var wg sync.WaitGroup
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var product Products
	json.Unmarshal(responseData, &product.Product)

	chCategory := make(chan string, len(product.Product))
	chTitle := make(chan string, len(product.Product))
	chPrice := make(chan interface{}, len(product.Product))

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(product.Product); i++ {
			product.mu.Lock()
			chTitle <- product.Product[i].Title
			chPrice <- product.Product[i].Price
			chCategory <- product.Product[i].Category
			product.mu.Unlock()
		}
	}()

	wg.Wait()
	for i := 0; i < len(product.Product); i++ {
		dataTitle := <-chTitle
		dataPrice := <-chPrice
		dataCategory := <-chCategory
		fmt.Printf(" ==========\n title: %s \n price: %v \n category: %s \n ==========\n", dataTitle, dataPrice, dataCategory)
	}

}
