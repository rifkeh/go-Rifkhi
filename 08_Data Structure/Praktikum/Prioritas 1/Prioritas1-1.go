package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(arrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(arrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(arrayMerge([]string{}, []string{}))
}

func arrayMerge(a, b []string) []string {
	c := make([]string, len(a))
	var isSame bool
	copy(c, a)
	for _, slice2 := range b {
		isSame = false
		for _, slice := range a {
			if slice2 == slice {
				isSame = true
			}
		}
		if isSame == false {
			c = append(c, slice2)
		}
	}
	return c
}
