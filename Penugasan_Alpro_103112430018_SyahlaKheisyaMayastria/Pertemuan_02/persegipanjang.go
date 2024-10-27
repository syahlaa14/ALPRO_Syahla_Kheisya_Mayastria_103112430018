package main

import (
	"fmt"
)

func main() {
	var panjang, lebar float64

	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scan(&lebar)

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Printf("Luas persegi panjang adalah: %.2f\n", luas)
	fmt.Printf("Keliling persegi panjang adalah: %.2f\n", keliling)
}
