package main
import "fmt"

func main() {

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)


	var counterX, counter0 int

	fmt.Println("Masukkan 9 bilangan (0 atau x):")
	for i := 0; i < 9; i++ {
		var input int
		fmt.Scan(&input)

		if input == x {
			counterX++
		} else if input == 0 {
			counter0++
		}
	}


	if counterX > counter0 {
		fmt.Println("Modus: ", x)
	} else if counterX < counter0 {
		fmt.Println("Modus: 0")
	} else {
		fmt.Println("Tidak ada modus yang jelas, kedua nilai muncul dengan frekuensi yang sama.")
	}
}

// MULAI
//   DEKLARASIKAN x, counterX, counter0 SEBAGAI INTEGER
//   SETEL counterX = 0
//   SETEL counter0 = 0

//   CETAK "Masukkan nilai x: "
//   INPUT x

//   CETAK "Masukkan 9 bilangan (0 atau x):"
//   UNTUK i DARI 0 SAMPAI 8 LAKUKAN:
//     DEKLARASIKAN input SEBAGAI INTEGER
//     INPUT input

//     JIKA input == x MAKA
//       counterX = counterX + 1
//     LAIN JIKA input == 0 MAKA
//       counter0 = counter0 + 1
//     AKHIR JIKA
//   AKHIR UNTUK

//   JIKA counterX > counter0 MAKA
//     CETAK "Modus: ", x
//   LAIN JIKA counterX < counter0 MAKA
//     CETAK "Modus: 0"
//   LAIN
//     CETAK "Tidak ada modus yang jelas, kedua nilai muncul dengan frekuensi yang sama."
//   AKHIR JIKA
// SELESAI