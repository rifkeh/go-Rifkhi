package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 0
	fmt.Printf("Masukkan banyaknya baris bintang: ")
	fmt.Scanln(&n)
	for i := n - 1; i >= 0; i-- {
		fmt.Printf(strings.Repeat(" ", i) + "*" + strings.Repeat(" *", n-(i+1)) + "\n")
	}

}

//      *
//     * *
//    * * *
//   * * * *
//  * * * * *
// * * * * * *
