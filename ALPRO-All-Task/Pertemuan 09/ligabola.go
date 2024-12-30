package main
import "fmt"


func main() {
	var gol1, gol2, gol3, gol4 int
	fmt.Print("Masukkan jumlah gol empat tim (dipisahkan dengan spasi): ")
	fmt.Scan(&gol1, &gol2, &gol3, &gol4)

	maksGol := gol1
	if gol2 > maksGol {
		maksGol = gol2
	}
	if gol3 > maksGol {
		maksGol = gol3
	}
	if gol4 > maksGol {
		maksGol = gol4
	}

	minGol := gol1
	if gol2 < minGol {
		minGol = gol2
	}
	if gol3 < minGol {
		minGol = gol3
	}
	if gol4 < minGol {
		minGol = gol4
	}

	fmt.Printf("Gol terbanyak: %d\nGol tersedikit: %d\n", maksGol, minGol)
}

// MULAI
//   DEKLARASIKAN gol1, gol2, gol3, gol4 SEBAGAI INTEGER
//   CETAK "Masukkan jumlah gol empat tim (dipisahkan dengan spasi): "
//   INPUT gol1, gol2, gol3, gol4

//   TETAPKAN maksGol = gol1
//   JIKA gol2 > maksGol MAKA
//     maksGol = gol2
//   AKHIR JIKA
//   JIKA gol3 > maksGol MAKA
//     maksGol = gol3
//   AKHIR JIKA
//   JIKA gol4 > maksGol MAKA
//     maksGol = gol4
//   AKHIR JIKA

//   TETAPKAN minGol = gol1
//   JIKA gol2 < minGol MAKA
//     minGol = gol2
//   AKHIR JIKA
//   JIKA gol3 < minGol MAKA
//     minGol = gol3
//   AKHIR JIKA
//   JIKA gol4 < minGol MAKA
//     minGol = gol4
//   AKHIR JIKA

//   CETAK "Gol terbanyak: ", maksGol
//   CETAK "Gol tersedikit: ", minGol
// SELESAI
