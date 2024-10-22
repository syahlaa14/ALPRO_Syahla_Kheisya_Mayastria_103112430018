package main

import (
	"fmt"
)

func terurutMengecil(n int) bool {
	// Ambil digit masing-masing
	digit1 := n / 100       // digit pertama
	digit2 := (n / 10) % 10 // digit kedua
	digit3 := n % 10        // digit ketiga

	// Cek apakah digit pertama lebih besar dari digit kedua
	// dan digit kedua lebih besar dari digit ketiga
	return digit1 > digit2 && digit2 > digit3
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan bulat 3 digit: ")
	fmt.Scan(&angka)

	// Panggil fungsi terurutMengecil dan cetak hasilnya
	hasil := terurutMengecil(angka)
	fmt.Println(hasil)
}
