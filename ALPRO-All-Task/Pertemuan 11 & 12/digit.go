package main
import "fmt"

func main() {

	var X int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&X)

	var angka []int
	total := 0

	for X > 0 {
		digit := X % 10  
		angka = append(angka, digit) 
		total += digit    
		X /= 10   
	}

	for i := 0; i < len(angka); i++ {
		fmt.Print(angka[i], " ")
	}
	fmt.Println() 

	fmt.Println(total)
}

// MULAI
//   DEKLARASIKAN X SEBAGAI INTEGER
//   DEKLARASIKAN angka SEBAGAI LIST INTEGER
//   DEKLARASIKAN total SEBAGAI INTEGER DENGAN NILAI AWAL 0

//   CETAK "Masukkan bilangan bulat positif: "
//   INPUT X

//   SELAMA X > 0 LAKUKAN:
//     digit = X MOD 10            // Ambil digit terakhir dari X
//     TAMBAHKAN digit KE angka    // Simpan digit ke dalam list angka
//     total = total + digit       // Tambahkan digit ke total
//     X = X DIV 10                // Hapus digit terakhir dari X
//   AKHIR SELAMA

//   UNTUK i DARI 0 HINGGA PANJANG(angka)-1 LAKUKAN:
//     CETAK angka[i] DI BARIS YANG SAMA
//   AKHIR UNTUK
//   CETAK BARIS BARU

//   CETAK total
// SELESAI