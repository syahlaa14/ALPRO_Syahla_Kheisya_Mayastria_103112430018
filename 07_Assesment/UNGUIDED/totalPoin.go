package main

import "fmt"

func main() {
    var jumlahBarang, totalPoin int
    fmt.Print("Masukkan jumlah barang yang dibeli: ")
    fmt.Scan(&jumlahBarang)

    for i := 1; i <= jumlahBarang; i++ {
        poin := 10 + 5*(i/6) 
        totalPoin += poin
    }

    fmt.Println("Total poin yang didapat:", totalPoin, "poin")
}
