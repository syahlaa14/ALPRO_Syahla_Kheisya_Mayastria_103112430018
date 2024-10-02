package main

import (
	"fmt"
)

func main() {

	var sisi int = 27        // Menentukan panjang sisi alun-alun
	var keliling, luas int

	keliling = 4 * sisi     // Rumus menghitung luas persegi
	luas = sisi * sisi      // Rumus menghitung keleling persegi

	fmt.Println("Keliling alun-alun adalah:", keliling)
	fmt.Println("Luas alun-alun adalah:", luas)
}
