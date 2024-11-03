package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	var isXFaktorY bool
	if y%x == 0 {
		isXFaktorY = true
	}
	if !isXFaktorY {
		isXFaktorY = false
	}

	var isYFaktorX bool
	if x%y == 0 {
		isYFaktorX = true
	}
	if !isYFaktorX {
		isYFaktorX = false
	}

	println(isXFaktorY)
	println(isYFaktorX)
}
