package main

import "fmt"

func main() {

	s := "abcdefg"
	sLen := len(s) - 1

	for i := range s {
		fmt.Print(string(s[sLen-i]))
	}

}
