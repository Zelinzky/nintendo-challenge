package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size int
	fmt.Scan(&size)

	// a is the input tab to encrypt
	// b is output tab
	var a, b = make([]int, size/16), make([]int, size/16)
	var astring string

	// read size/16 integers to a
	for i := 0; i < size/16; i++ {
		fmt.Scan(&astring)
		middle, _ := strconv.ParseInt(astring, 16, 64)
		a[i] = int(middle)
	}

	// here they initialize the int array to 0, but golang has 0 value by default

	// Main encryption
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// MAGIC!!
			// unaltered magic
			b[(i+j)/32] ^= ((a[i/32] >> (i % 32)) & (a[j/32+size/32] >> (j % 32)) & 1) << ((i + j) % 32)
		}
	}

	// Printing encoded message
	for i := 0; i < size/16; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%8x", b[i])
	}
}
