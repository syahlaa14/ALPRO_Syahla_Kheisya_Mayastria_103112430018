package main

import "fmt"

func main() {
	// Deklarasi variabel
	var angka, digit, jumlahDigitGenap int

	// Meminta pengguna memasukkan angka
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Inisialisasi jumlah digit genap
	jumlahDigitGenap = 0

	// Perulangan untuk memproses setiap digit
	for angka > 0 {
		digit = angka % 10 // Mengambil digit terakhir
		if digit%2 == 0 && digit != 0 { // Mengecek apakah digit genap dan bukan nol
			jumlahDigitGenap++
		}
		angka = angka / 10 // Menghapus digit terakhir
	}

	// Menampilkan jumlah digit genap
	fmt.Println("Jumlah digit genap (tidak termasuk nol):", jumlahDigitGenap)
}
