package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int

	fmt.Print("pilih jenis kendaraan (motor, mobil, truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	var tarifPerJam int

	switch kendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("jenis kendaraan tidak valid!")
		return
	}

	totalBiaya := tarifPerJam * durasi
	fmt.Printf("Rp.%d\n", totalBiaya)
}
