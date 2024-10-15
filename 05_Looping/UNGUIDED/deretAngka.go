package main

import (
	"fmt"
)

func main() {
	var n, hasil int

	fmt.Print("Masukkan bilangan bulat positif: ")		// User inputkan bilangan positif
	fmt.Scan(&n)

	hasil = n * (n + 1) / 2	// Rumus menghitung deret angka 

    fmt.Println("Jumlah dari deret 1 hingga", n, "adalah", hasil)
}
