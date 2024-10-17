package main

import "fmt"

func main() {
	var n int		// variabel untuk input
	fmt.Println("Masukkan bilangan bulat:")
	fmt.Scan(&n)

	hasil := 1

	for i := 1; i <= n; i++ {
		hasil *= i
	}

	fmt.Println("Hasil faktorial dari", n, "adalah:", hasil)
}
