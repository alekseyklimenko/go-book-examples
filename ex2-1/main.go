package main

import (
	"fmt"
	"github.com/alekseyklimenko/go-book-examples/conv"
)

func main() {
	fmt.Printf("%g\n", conv.BoilingC-conv.FreezingC)
	boilingK := conv.CToK(conv.BoilingC)
	fmt.Printf("%g\n", boilingK)
}
