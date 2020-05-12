package main

import "fmt"

var pc [256]byte

func main() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	var val uint64 = 22010000
	fmt.Printf("population count for %d = %d\n", val, popCount(val))
	fmt.Printf("population count for %d = %d\n", val, popCount2(val))
	fmt.Printf("population count for %d = %d\n", val, popCount3(val))
	fmt.Printf("population count for %d = %d\n", val, popCount4(val))
}

// PopCount returns the population count (number of set bits) of x.
func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCount2(x uint64) int {
	var res byte
	for i := 0; i < 8; i++ {
		res += pc[byte(x>>(i*8))]
	}
	return int(res)
}

func popCount3(x uint64) int {
	var res uint64
	for i := 0; i < 64; i++ {
		res += (x >> i) & 1
	}
	return int(res)
}

func popCount4(x uint64) int {
	res := 0
	for x > 0 {
		res++
		x = x & (x - 1)
	}
	return res
}
