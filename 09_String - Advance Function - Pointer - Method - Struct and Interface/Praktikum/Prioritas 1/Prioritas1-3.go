package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if strings.Compare(a, b) == 1 {
		return b
	} else if strings.Compare(a, b) == -1 {
		return a
	} else {
		if a == "" || b == "" {
			return ""
		} else {
			return a
		}
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA

}
