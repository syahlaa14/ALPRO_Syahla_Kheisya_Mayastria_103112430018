package main

import (
	"fmt"
)

func main() {
	const pi = 3.1415926535 
	var r int

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&r)

	volume := (4.0 / 3.0) * pi * float64(r*r*r) 	// Menghitung volume bola tanpa 

	luas := 4 * pi * float64(r*r)		// Menghitung luas permukaan bola 

	hasil := fmt.Sprintf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
	fmt.Println(hasil)
}
