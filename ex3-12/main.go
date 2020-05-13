package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s isAnagram-%v\n", os.Args[i], isAnagram(os.Args[i]))
	}
}

func isAnagram(s string) bool {
	length := len(s)
	n := length / 2
	for i := 0; i <= n; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
