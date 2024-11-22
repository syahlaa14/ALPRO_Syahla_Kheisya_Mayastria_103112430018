package main

import "fmt"

func main() {
	var input int

	fmt.Scan(&input)

	switch {
	case input%10 == 0:
		fmt.Print("Kategori: Bilangan Kelipatan 10\n")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", input, input/10)
	case input%5 == 0 && input >5:
		fmt.Print("Kategori: Bilangan Kelipatan 5\n")
		fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d\n", input, input*input)
	case input%2 == 0:
		fmt.Print("Kategori: Bilangan Genap\n")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", input, input+1, input*(input+1))
	default:
		fmt.Print("Kategori: Bilangan Ganjil\n")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", input, input+1, input+(input+1))
	}
}