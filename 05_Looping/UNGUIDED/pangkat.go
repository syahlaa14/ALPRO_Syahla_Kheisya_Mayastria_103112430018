package main

import "fmt"

func main() {
	var bilangan, pangkat int		// variabel untuk input
	fmt.Println("Masukkan bilangan:")
	fmt.Scan(&bilangan)
	fmt.Println("Masukkan bilangan pangkat:")
	fmt.Scan(&pangkat)

	hasil := 1

	for i := 0; i < pangkat; i++ {
		hasil *= bilangan
	}

	fmt.Println("Hasil dari", bilangan, "dipangkatkan", pangkat, "adalah:", hasil)
}
