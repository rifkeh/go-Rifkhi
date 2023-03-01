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
	nameEncode := s.name
	var alphabet, key string = "abcdefghijklmnopqrstuvwxyz ", "THELADYISMRGNWTFXZOQVBJCKP "
	for i := 0; i <= len(alphabet)-2; i++ {
		nameEncode = strings.Replace(nameEncode, alphabet[i:i+1], key[i:i+1], -1)
	}
	nameEncode = strings.ToLower(nameEncode)
	return nameEncode
}

func (s *student) Decode() string {
	nameEncode := s.name
	var alphabet, key string = "theladyismrgnwtfxzoqvbjckp ", "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	for i := 0; i <= len(alphabet)-2; i++ {
		nameEncode = strings.Replace(nameEncode, alphabet[i:i+1], key[i:i+1], -1)
	}
	nameEncode = strings.ToLower(nameEncode)
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
