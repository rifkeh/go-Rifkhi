package main

import (
	"fmt"
	"strconv"
)

func solveXYZ(A, B, C int) (r1, r2, r3 string) {
	isFound := false
	for x := -100; x <= 100; x++ {
		for y := -100; y <= 100; y++ {
			z := A - x - y
			if x*y*z == B && x*x+y*y+z*z == C {
				r1, r2, r3 = strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z)
				isFound = true
			}
			if isFound == true {
				break
			}
		}
		if isFound == true {
			break
		}
	}
	if isFound == false {
		return "tidak ditemukan", "", ""
	}
	return r1, r2, r3
}
func main() {
	fmt.Println(solveXYZ(1, 2, 3))
	fmt.Println(solveXYZ(6, 6, 14))
}
