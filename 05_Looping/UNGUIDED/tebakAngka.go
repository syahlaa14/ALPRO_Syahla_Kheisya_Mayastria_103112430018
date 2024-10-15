package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	angkaAcak := rand.Intn(100) + 1
	var tebakan, kesempatan int

	kesempatan = 5

	fmt.Println("Coba tebak angka dari 1 hingga 100.")		// Print perintah 
	fmt.Println("Ada 5 kesempatan untuk menebak")

	for i := 1; i <= kesempatan; i++ {
		fmt.Printf("\nCoba %d: Masukkan angka tebakan: ", i)		// User masukkan angka
		fmt.Scan(&tebakan)

		if tebakan == angkaAcak {
			fmt.Println("Tebakanmu benar!")
			return 
		} else if tebakan < angkaAcak {
			fmt.Println("Angka terlalu kecil!")
		} else {
			fmt.Println("Angka terlalu besar!")
		}
	}

	fmt.Printf("\nKesempatan telah habis. Angka yang benar adalah %d.\n", angkaAcak)
}
