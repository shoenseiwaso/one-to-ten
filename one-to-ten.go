package main

import (
	"fmt"
)

func main() {
	z := 1
	fmt.Printf("%v enzo\n", z)
	for ; z < 100000000; z = z + 1 {
	}
	fmt.Printf("%v enzo\n", z)
}