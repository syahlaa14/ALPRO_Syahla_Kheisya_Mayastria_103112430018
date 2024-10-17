package main

import (
	"fmt"
	"math"
)

func main() {

	var jmlKerucut int
	fmt.Scan(&jmlKerucut)

	var jari, tinggi float64		// variabel untuk menyimpan jari-jari dan tinggi

	for i := 0; i < jmlKerucut; i++ {		// untuk menghitung setiap kerucut
		fmt.Printf("Masukkan jari-jari: ")
		fmt.Scan(&jari)
		
		fmt.Printf("Masukkan tinggi: ")
		fmt.Scan(&tinggi)

		volume := (1.0 / 3.0) * math.Pi * jari * jari * tinggi		// menghitung volume
		fmt.Println("Volume kerucut adalah:", volume)
	}
}
