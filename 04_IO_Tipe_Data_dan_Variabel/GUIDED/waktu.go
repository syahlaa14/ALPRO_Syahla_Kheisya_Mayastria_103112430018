package main

import (
	"fmt"
)

func main() {
	var totalDetik int							// Input jumlah detik dari user
	fmt.Print("Masukkan jumlah detik: ")
	fmt.Scan(&totalDetik)

	var jam, menit, detik int		// Menghitung jam, menit, dan sisa detik
	jam = totalDetik / 3600
	menit = (totalDetik % 3600) / 60
	detik = totalDetik % 60

	fmt.Print(totalDetik, " detik = ")
	fmt.Print(jam, "jam ")
	fmt.Print(menit, "menit ")
	fmt.Println(detik, "detik")
}
