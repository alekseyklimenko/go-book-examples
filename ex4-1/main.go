package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) < 2 {
		log.Fatal("Need two parameters for calculation sha256 population diff\n")
	}
	sha1 := sha256.Sum256([]byte(os.Args[1]))
	fmt.Printf("sha1: %x\n", sha1)
	sha2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("sha2: %x\n", sha2)
	diff := popDiff(sha1, sha2)
	fmt.Printf("Bits diff: %d\n", diff)
}

func popDiff(sha1, sha2 [32]byte) int {
	res := 0
	for i, val := range sha1 {
		for j := 0; j < 8; j++ {
			if ((val >> j) & 1) != ((sha2[i] >> j) & 1) {
				res++
			}
		}
	}
	return res
}
