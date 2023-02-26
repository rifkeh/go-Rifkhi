package main

import "fmt"

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}

func Mapping(slice []string) map[string]int {
	var banyak int
	hasil := make(map[string]int)
	var isSame bool
	for i, temp := range slice {
		banyak = 0
		isSame = false
		for j, temp1 := range slice {
			if i < j {
				if temp == temp1 {
					isSame = true
					break
				}
			} else {
				if temp == temp1 {
					banyak += 1
				}
			}
		}
		if isSame == false {
			hasil[temp] = banyak
		}
	}
	return hasil
}
