package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/alekseyklimenko/go-book-examples/conv"
	"os"
	"strconv"
)

var convType = flag.String("t", "", "Type of convert (\"m2f\", \"f2m\", \"k2l\", \"l2k\")")
var value = flag.Float64("val", 0, "Value to convert")

func main() {
	flag.Parse()
	var val float64
	if *value == 0 {
		fmt.Print("Enter a value to convert: ")
		input := bufio.NewReader(os.Stdin)
		strval, err := input.ReadString('\n')
		if err != nil {
			intval, err := strconv.Atoi(strval)
			if err != nil {
				val = float64(intval)
			}
		}
	} else {
		val = *value
	}
	switch *convType {
	case "m2f":
		result := conv.MeterToFt(conv.Meter(val))
		fmt.Printf("result=%s\n", result)
	case "f2m":
		result := conv.FtToMeter(conv.Ft(val))
		fmt.Printf("result=%s\n", result)
	case "k2l":
		result := conv.KgToLb(conv.Kg(val))
		fmt.Printf("result=%s\n", result)
	case "l2k":
		result := conv.LbToKg(conv.Lb(val))
		fmt.Printf("result=%s\n", result)
	}
}
