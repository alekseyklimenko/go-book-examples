package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var shaBits = flag.Int("sha", 256, "Bits in sha digest (256, 384 or 512)")
var value = flag.String("val", "", "Value to calculate sha")

func main() {
	flag.Parse()
	var res []byte
	switch *shaBits {
	case 256:
		sha := sha256.Sum256([]byte(*value))
		res = sha[:32]
	case 384:
		sha := sha512.Sum384([]byte(*value))
		res = sha[:48]
	case 512:
		sha := sha512.Sum512([]byte(*value))
		res = sha[:64]
	}
	fmt.Printf("Calculated digest: %x\n", res)
}
