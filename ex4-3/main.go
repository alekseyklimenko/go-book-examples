package main

import (
	"fmt"
	"unicode"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	//	reverse(a[:])
	reverse2(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2]) //[1 0 2 3 4 5]
	reverse(s[2:]) //[1 0 5 4 3 2]
	reverse(s)
	fmt.Println(s) //[2 3 4 5 0 1]

	s2 := []int{0, 1, 2, 3, 4, 5}
	s2 = rotate(s2, 2)
	fmt.Println(s2) //[2 3 4 5 0 1]
	//!-slice

	data := []string{"one", "two", "two", "two", "three", "three"}
	data = dedupl(data)
	fmt.Println(data)

	str := "Hello,      世界"
	r := []rune(str)
	r = utfSpaceDedup(r)
	fmt.Println(r)
	fmt.Println(string(r))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func rotate(s []int, len int) []int {
	res := s[len:]
	res = append(res, s[:len]...)
	return res
}

func dedupl(strings []string) []string {
	newCur := 1
	var pPrevStr = &strings[0]
	for i := 1; i < len(strings); i++ {
		if strings[i] != *pPrevStr {
			strings[newCur] = strings[i]
			newCur++
			pPrevStr = &strings[i]
		}
	}
	return strings[:newCur]
}

func utfSpaceDedup(str []rune) []rune {
	curIndex := 0
	findSpace := false
	for i := 0; i < len(str); i++ {
		if unicode.IsSpace(str[i]) {
			if !findSpace {
				findSpace = true
			}
		} else {
			if findSpace {
				str[curIndex] = ' '
				curIndex++
				findSpace = false
			}
			str[curIndex] = str[i]
			curIndex++
		}
	}
	return str[:curIndex]
}
