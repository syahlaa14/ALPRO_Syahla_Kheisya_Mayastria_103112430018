package main
import "fmt"

func main() {

	var T int
	fmt.Print("Masukkan kapasitas tanki (T): ")
	fmt.Scan(&T)

	var totalAir int

	for {
		var V int
		fmt.Print("Masukkan volume ember (V) atau 0 untuk berhenti: ")
		fmt.Scan(&V)

		if V == 0 {
			break
		}

		totalAir += V

		if totalAir >= T {
			fmt.Println(true) 
		} else {
			fmt.Println(false) 
		}
	}
}

// MULAI
//   DEKLARASIKAN T SEBAGAI INTEGER
//   CETAK "Masukkan kapasitas tanki (T): "
//   INPUT T

//   DEKLARASIKAN totalAir SEBAGAI INTEGER DENGAN NILAI AWAL 0

//   ULANGI SELAMANYA:
//     DEKLARASIKAN V SEBAGAI INTEGER
//     CETAK "Masukkan volume ember (V) atau 0 untuk berhenti: "
//     INPUT V

//     JIKA V == 0 MAKA
//       HENTIKAN ULANGAN
//     AKHIR JIKA

//     totalAir = totalAir + V

//     JIKA totalAir >= T MAKA
//       CETAK true
//     LAIN
//       CETAK false
//     AKHIR JIKA
//   AKHIR ULANGI
// SELESAI