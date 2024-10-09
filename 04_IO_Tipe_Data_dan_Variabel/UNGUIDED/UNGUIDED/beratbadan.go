package main

import (
	"fmt"
)

func main() {
	var bmi, tinggi, berat float64
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggi)

	berat = bmi * (tinggi * tinggi)		// Menghitung berat badan
	fmt.Printf("Berat badan seseorang adalah: %.2f kg\n", berat)
}
