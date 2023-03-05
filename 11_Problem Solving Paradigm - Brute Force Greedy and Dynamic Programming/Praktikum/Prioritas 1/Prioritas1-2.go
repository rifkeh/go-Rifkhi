package main

import "fmt"

func pascalTriangle(numRows int) (result [][]int) {

	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, temp)
	}
	return result
}

func main() {
	fmt.Println(pascalTriangle(5))
}
