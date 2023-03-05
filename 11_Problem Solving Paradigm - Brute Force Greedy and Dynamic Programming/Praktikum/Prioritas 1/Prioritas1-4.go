package main

import "fmt"

func solveXYZ(A, B, C int) (r1, r2, r3 int) {
	isFound := false
	for x := -100; x <= 100; x++ {
		for y := -100; y <= 100; y++ {
			z := A - x - y
			if x*y*z == B && x*x+y*y+z*z == C {
				r1, r2, r3 = x, y, z
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
		fmt.Println("tidak ditemukan")
		return
	}
	return r1, r2, r3
}
func main() {
	fmt.Println(solveXYZ(1, 2, 3))
	fmt.Println(solveXYZ(6, 6, 14))
}
