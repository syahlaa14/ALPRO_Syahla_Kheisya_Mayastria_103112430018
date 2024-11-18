package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih dari 1.")
		return
	}

	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}

	fmt.Print("Faktor: ")
	for _, f := range faktor {
		fmt.Print(f, " ")
	}
	fmt.Println()

	isPrima := len(faktor) == 2
	fmt.Println("Prima:", isPrima)
}
