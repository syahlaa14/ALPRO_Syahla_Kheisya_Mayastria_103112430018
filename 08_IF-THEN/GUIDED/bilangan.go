package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	if angka >= 0 {
		fmt.Println("positif")
	}
	if angka < 0 {
		fmt.Println("bukan positif")
	}
}
