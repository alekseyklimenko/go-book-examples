package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	steps := n / 3
	buf.WriteString(s[:n-steps*3])
	for i := steps; i > 0; i-- {
		j := n - i*3
		buf.WriteString("," + s[j:j+3])
	}
	return buf.String()
}
