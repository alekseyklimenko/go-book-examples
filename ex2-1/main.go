package main

import (
	"fmt"
	"github.com/alekseyklimenko/go-book-examples/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingK := tempconv.CToK(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingK)
}
