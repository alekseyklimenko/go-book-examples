package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[strings.ToLower(input.Text())]++
	}
	for word, n := range counts {
		fmt.Printf("%d\t%s\n", n, word)
	}
}
