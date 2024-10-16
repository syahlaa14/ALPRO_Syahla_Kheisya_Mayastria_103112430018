package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukan angka 1: ")
	fmt.Scan(&a)
	fmt.Print("Masukan angka 2: ")
	fmt.Scan(&b)
	
	hasil := 0
	for i := 1; i <= b; i+=1 {
		hasil = hasil + a 
	}
	fmt.Println("Hasil perkalian adalah:", hasil)
}