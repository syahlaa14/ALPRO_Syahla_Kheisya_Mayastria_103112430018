package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Println("Masukkan bilangan bulat:")
    fmt.Scanln(&x)

	if x == -5 {
		fmt.Println("Nilai x tidak boleh -5, karena menyebabkan pembagian dengan nol.")
	} else {
		result := (2.0 / float64(x+5)) + 5.0
		fmt.Printf("Hasil dari f(x) = (2 / (x + 5)) + 5 adalah: %.2f\n", result)
	}
}
