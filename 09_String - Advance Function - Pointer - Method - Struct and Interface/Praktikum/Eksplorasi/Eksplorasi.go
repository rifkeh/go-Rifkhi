package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	nameEncode := ""
	var alphabet, key string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ", "theladyismrgnwtfxzoqvbjckpTHELADYISMRGNWTFXZOQVBJCKP "
	for i := 0; i <= len(s.name)-1; i++ {
		for j := 0; j <= len(alphabet)-2; j++ {
			if strings.Contains(s.name[i:i+1], alphabet[j:j+1]) {
				nameEncode += strings.Replace(s.name[i:i+1], alphabet[j:j+1], key[j:j+1], -1)
			}
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	nameEncode := ""
	var alphabet, key string = "theladyismrgnwtfxzoqvbjckpTHELADYISMRGNWTFXZOQVBJCKP ", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	for i := 0; i <= len(s.name)-1; i++ {
		for j := 0; j <= len(alphabet)-2; j++ {
			if strings.Contains(s.name[i:i+1], alphabet[j:j+1]) {
				nameEncode += strings.Replace(s.name[i:i+1], alphabet[j:j+1], key[j:j+1], -1)
			}
		}
	}
	return nameEncode
}

func main() {
	var menu int
	a := student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	}
}
