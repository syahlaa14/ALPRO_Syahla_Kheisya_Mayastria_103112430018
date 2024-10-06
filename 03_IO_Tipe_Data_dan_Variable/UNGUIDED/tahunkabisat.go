package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Masukkan tahun: ")			// Masukan input dari user
	fmt.Scanln(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
