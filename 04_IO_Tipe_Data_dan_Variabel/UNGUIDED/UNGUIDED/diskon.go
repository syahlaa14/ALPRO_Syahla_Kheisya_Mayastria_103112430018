package main

import (
	"fmt"
)

func main() {
	var totalBelanja, diskonPersen int

	fmt.Print("Masukkan total belanja awal: ")		
	fmt.Scan(&totalBelanja)

	fmt.Print("Masukkan besaran diskon (dalam persen): ")
	fmt.Scan(&diskonPersen)

	diskon := float64(diskonPersen) / 100 * float64(totalBelanja)
	totalSetelahDiskon := float64(totalBelanja) - diskon
	
	fmt.Println("Total harga setelah diskon adalah:", totalSetelahDiskon)
}
