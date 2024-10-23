package main

import "fmt"

func main() {
    var angka1, angka2, total int       //deklarasi variabel

    fmt.Scan(&angka1)
    fmt.Scan(&angka2)

    for i := angka1; i <= angka2; i++ {     // penjumlahan dari angka1 hingga angka2
        total += i
    }

    fmt.Println("Hasil penjumlahan dari", angka1, "sampai", angka2, "adalah:", total)
}
