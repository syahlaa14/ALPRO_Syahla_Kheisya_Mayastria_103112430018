package main

import "fmt"

func main() {
    var dinar, dirham, fals, qirat, input, sisa input   // deklarasi variabel 

    fmt.Print("Masukkan jumlah qirat: ")        // input dari user
    fmt.Scan(&input)

    dinar = input / (10 * 10 * 6)       // rumus menghitung penukaran uang
    sisa = input % (10 * 10 * 6)

    dirham = sisa / (10 * 6)
    sisa = sisa % (10 * 6)

    fals = sisa / 6
    qirat = sisa % 6

    fmt.Println("Hasil penukaran adalah : ")
    fmt.Println("Dinar :", dinar)
    fmt.Println("Dirham :", dirham)
    fmt.Println("Fals :", fals)
    fmt.Println("Qirat :", qirat)
}

