package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var line, lineIndex string
	counts := make(map[string]int)
	fileNames := make(map[string]int)
	files := os.Args[1:]
	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			input := bufio.NewScanner(f)
			for input.Scan() {
				line = input.Text()
				lineIndex = arg + line
				counts[lineIndex]++
				if counts[lineIndex] > 1 {
					fileNames[arg]++
				}
			}
			f.Close()
		}
	}
	for name, _ := range fileNames {
		fmt.Printf("%s\n", name)
	}
}
