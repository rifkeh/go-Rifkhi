package main

import "fmt"

type pair struct {
	name string
	idx  int
}

// Supaya bisa dimasukkan kedalam parameter untuk melakukan sorting
type pairs []pair

// Quicksort untuk melakukan pengurutan pada slice struct nanti dimana time complexity O(n logn)
func (t pairs) quickSort() {
	if len(t) <= 1 {
		return
	}
	pivot := len(t) - 1
	i := 0
	for j := 0; j < pivot; j++ {
		if t[j].idx < t[pivot].idx {
			t[i], t[j] = t[j], t[i]
			i++
		}
	}
	t[i], t[pivot] = t[pivot], t[i]
	pairs(t[:i]).quickSort()
	pairs(t[i+1:]).quickSort()
}

func mostAppearItem(items []string) []pair {
	// Membuat map sementara untuk menghitung item yang sama dengan menghasilkan key dan value yang berpadanan
	tempMap := make(map[string]int)

	for _, item := range items {
		tempMap[item] += 1
	}

	//Deklarasi pembuatan slice of struct mengacu pada struct pair
	tempSlice := make(pairs, len(tempMap))

	//Assign key dan value dari Map kedalam slice of struct
	i := 0
	for name, idx := range tempMap {
		tempSlice[i] = pair{name, idx}
		i += 1
	}

	// Lakukan pengurutan bedarsarkan value map atau idx dari slice
	tempSlice.quickSort()

	return tempSlice
}

func main() {
	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(mostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(mostAppearItem([]string{"football", "basketball", "tenis"}))
}
