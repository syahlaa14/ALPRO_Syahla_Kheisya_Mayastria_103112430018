package main

import (
    "fmt"
)

func main() {
    var angka1, angka2 float64
    var operator string

    fmt.Print("Masukkan perhitungan  ( -, +, *, / ): ")    // Masukan inputan dari user
    fmt.Scan(&angka1, &operator, &angka2)

    switch operator {
    case "+":               // Penjumlahan
        fmt.Println("Hasil:", angka1, "+", angka2, "=", angka1+angka2)
    case "-":               // Pengurangan
        fmt.Println("Hasil:", angka1, "-", angka2, "=", angka1-angka2)
    case "*":               // Perkalian
        fmt.Println("Hasil:", angka1, "*", angka2, "=", angka1*angka2)
    case "/":               // Pembagian
        if angka2 != 0 {    // Tdk bisa dibagi dengan 0
            fmt.Println("Hasil:", angka1, "/", angka2, "=", angka1/angka2)
        } else {
            fmt.Println("Tidak boleh dibagi dengan 0.")
        }
    default:    // Operator salah 
        fmt.Println("Operator tidak sesuai, silakan diulang kembali.")
    }
}
