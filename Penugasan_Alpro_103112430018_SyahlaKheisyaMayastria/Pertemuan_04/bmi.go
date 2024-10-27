package main

import (
	"fmt"
)

func main() {
	var bmi, berat, tinggi float64
	var nama string

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&berat)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)

	fmt.Print("Informasi BMI: ")
	fmt.Println("Nama: ", nama)
	fmt.Printf("Berat: %.2f\n", berat)
	fmt.Printf("Tinggi: %.2f\n", tinggi)
	fmt.Printf("BMI: %.2f\n", bmi)
}
