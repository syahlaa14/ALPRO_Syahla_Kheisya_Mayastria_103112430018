package main

import (
	"fmt"
)

func main() {
	var angka int		// Deklarasi variabel
	angka = 2

	for angka <= 50 {
		fmt.Println(angka)
		angka += 2
	}
}
