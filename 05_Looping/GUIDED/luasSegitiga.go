package main

import "fmt"

func main() {
	var n, a, t int

	fmt.Println("Berapa kali anda ingin menghitung: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("Masukan alas segitiga: ")
		fmt.Scan(&a)
		fmt.Print("Masukan tinggi segitiga: ")
		fmt.Scan(&t)
		p := a * t / 2
		fmt.Println("Hasilnya adalah:", p)
	}
}