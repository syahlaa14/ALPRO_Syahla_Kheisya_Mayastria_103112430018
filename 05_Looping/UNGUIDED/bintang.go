package main

import (
	"fmt"
)

func main() {
	var n, i, j  int
	fmt.Print("Masukkan jumlah baris: ")		// User inputkan jumlah baris yang dimau
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {			
		for j = 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
