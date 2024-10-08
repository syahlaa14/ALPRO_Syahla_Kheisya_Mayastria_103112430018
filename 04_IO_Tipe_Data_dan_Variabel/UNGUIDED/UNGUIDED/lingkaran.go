package main

import (
	"fmt"
)

func main() {

	var r, luas, keliling float64
	const pi = 22/7 

	fmt.Print("Masukkan jari-jari lingkaran: ")			// Memasukkan input
	fmt.Scanln(&r)

	luas = pi * r * r                  // Menghitung luas dan keliling lingkaran
	keliling = 2 * pi * r              

	fmt.Println("Luas lingkaran adalah :", luas)
	fmt.Println("Keliling lingkaran adalah :", keliling)
}
