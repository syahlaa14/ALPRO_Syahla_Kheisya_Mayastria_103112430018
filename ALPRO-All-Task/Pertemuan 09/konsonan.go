package main

import (
	"fmt"
	"strings"
)

func main() {
	var karakter string
	fmt.Scan(&karakter)

	karakter = strings.ToLower(karakter)

	vokal := "aeiou"

	if len(karakter) == 1 && karakter[0] >= 'a' && karakter[0] <= 'z' {
		if strings.Contains(vokal, karakter) {
			fmt.Println("bukan konsonan")
		} else {
			fmt.Println("konsonan")
		}
	} else {
		fmt.Println("bukan konsonan")
	}
}

// MULAI
//   DEKLARASIKAN karakter SEBAGAI STRING
//   INPUT karakter

//   UBAH karakter MENJADI huruf kecil (lowercase)

//   TETAPKAN vokal MENJADI "aeiou"

//   JIKA panjang karakter ADALAH 1 DAN karakter ADALAH huruf (a-z) MAKA
//     JIKA karakter ADA di dalam vokal MAKA
//       CETAK "bukan konsonan"
//     SELAIN ITU
//       CETAK "konsonan"
//     AKHIR JIKA
//   SELAIN ITU
//     CETAK "bukan konsonan"
//   AKHIR JIKA
// SELESAI
