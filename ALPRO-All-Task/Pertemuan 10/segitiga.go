package main
import "fmt"

func main() {

	var a, b, c int
	fmt.Print("Masukkan panjang sisi pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan panjang sisi kedua: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan panjang sisi ketiga: ")
	fmt.Scan(&c)

	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Panjang sisi harus bilangan positif")
		return
	}

	if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || b == c || a == c {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}

// MULAI
//   DEKLARASIKAN a, b, c SEBAGAI INTEGER
//   CETAK "Masukkan panjang sisi pertama: "
//   INPUT a
//   CETAK "Masukkan panjang sisi kedua: "
//   INPUT b
//   CETAK "Masukkan panjang sisi ketiga: "
//   INPUT c

//   JIKA a <= 0 ATAU b <= 0 ATAU c <= 0 MAKA
//     CETAK "Panjang sisi harus bilangan positif"
//     KELUAR DARI PROGRAM
//   AKHIR JIKA

//   JIKA a == b DAN b == c MAKA
//     CETAK "Segitiga Sama Sisi"
//   LAIN JIKA a == b ATAU b == c ATAU a == c MAKA
//     CETAK "Segitiga Sama Kaki"
//   LAIN
//     CETAK "Segitiga Sembarang"
//   AKHIR JIKA
// SELESAI