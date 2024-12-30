package main

import "fmt"


func main() {
	var jamMulai, menitMulai, jamSelesai, menitSelesai int
	fmt.Print("Masukkan waktu mulai parkir (jam menit): ")
	fmt.Scan(&jamMulai, &menitMulai)
	fmt.Print("Masukkan waktu selesai parkir (jam menit): ")
	fmt.Scan(&jamSelesai, &menitSelesai)

	menitMulaiTotal := (jamMulai % 12) * 60 + menitMulai
	menitSelesaiTotal := (jamSelesai % 12) * 60 + menitSelesai

	durasiMenit := menitSelesaiTotal - menitMulaiTotal
	if durasiMenit < 0 {
	}

	jumlahJam := durasiMenit / 60 
	jumlahMenit := durasiMenit % 60 

	fmt.Printf("Durasi parkir: %d jam %d menit\n", jumlahJam, jumlahMenit)
}

// MULAI
//   DEKLARASIKAN jamMulai, menitMulai, jamSelesai, menitSelesai SEBAGAI INTEGER
//   CETAK "Masukkan waktu mulai parkir (jam menit): "
//   INPUT jamMulai, menitMulai
//   CETAK "Masukkan waktu selesai parkir (jam menit): "
//   INPUT jamSelesai, menitSelesai

//   HITUNG menitMulaiTotal = (jamMulai MOD 12) * 60 + menitMulai
//   HITUNG menitSelesaiTotal = (jamSelesai MOD 12) * 60 + menitSelesai

//   HITUNG durasiMenit = menitSelesaiTotal - menitMulaiTotal
//   JIKA durasiMenit < 0 MAKA
//     DURASI MENIT DIADJUST UNTUK MENANGANI KASUS DURASI NEGATIF (BELUM DIISI DALAM KODE)
//   AKHIR JIKA

//   HITUNG jumlahJam = durasiMenit DIV 60
//   HITUNG jumlahMenit = durasiMenit MOD 60

//   CETAK "Durasi parkir: ", jumlahJam, " jam ", jumlahMenit, " menit"
// SELESAI