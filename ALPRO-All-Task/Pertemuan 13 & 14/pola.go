package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n) 

    if n%2 == 0 {
        fmt.Println("Input harus bilangan ganjil!")
        return
    }

    for i := 0; i < n; i++ { 
        for j := 0; j < n; j++ { 
            if j == i || j == n-i-1 {
                fmt.Print(i + 1)
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println() 
    }
}

// MULAI
//   DEKLARASIKAN n SEBAGAI INTEGER
//   INPUT n

//   JIKA n MOD 2 == 0 MAKA
//     CETAK "Input harus bilangan ganjil!"
//     KELUAR DARI PROGRAM
//   AKHIR JIKA

//   UNTUK i DARI 0 SAMPAI n-1 LAKUKAN:
//     UNTUK j DARI 0 SAMPAI n-1 LAKUKAN:
//       JIKA j == i ATAU j == n - i - 1 MAKA
//         CETAK i + 1 TANPA BARIS BARU
//       LAIN
//         CETAK " " TANPA BARIS BARU
//       AKHIR JIKA
//     AKHIR UNTUK
//     CETAK BARIS BARU
//   AKHIR UNTUK
// SELESAI