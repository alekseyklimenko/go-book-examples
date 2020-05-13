package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	startPos := 0
	if strings.HasPrefix(s, "-") {
		buf.WriteByte('-')
		startPos = 1
	}
	n := len(s) - startPos
	if n <= 3 {
		return s
	}
	steps := n / 3
	buf.WriteString(s[startPos : n-steps*3])
	for i := steps; i > 0; i-- {
		j := n - i*3
		if j != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[j : j+3])
	}
	return buf.String()
}
