package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan b: ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Nilai a harus lebih kecil atau sama dengan b.")
		return
	}

	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println() 
}