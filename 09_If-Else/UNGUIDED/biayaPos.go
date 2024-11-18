package main

import (
	"fmt"
)

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisaGram := berat % 1000
	biayaPerKg := 10000
	biayaTambahan := 0

	if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else {
		biayaTambahan = sisaGram * 15
	}

	if kg > 10 {
		biayaTambahan = 0
	}

	totalBiaya := (kg * biayaPerKg) + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", kg*biayaPerKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
