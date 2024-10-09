package main

import (
	"fmt"
)

func main() {
	var bmi, berat, tinggi float64			// Deklarasi variabel
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&berat)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)			// Menghitung BMI

	fmt.Printf("BMI Anda adalah: %.2f\n", bmi)
}
