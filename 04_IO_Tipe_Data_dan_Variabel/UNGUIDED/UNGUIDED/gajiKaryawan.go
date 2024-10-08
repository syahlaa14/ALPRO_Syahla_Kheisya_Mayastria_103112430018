package main

import (
    "fmt"
)

func main() {
    var jamPerMinggu, upahPerJam float64
    var jamLembur, jamNormal, totalGajiBulanan float64

    // Input jumlah jam kerja per minggu dan upah per jam
    fmt.Print("Masukkan jumlah jam kerja per minggu: ")
    fmt.Scan(&jamPerMinggu)
    fmt.Print("Masukkan upah per jam: ")
    fmt.Scan(&upahPerJam)

    // Jika jam kerja lebih dari 40 jam per minggu
    if jamPerMinggu > 40 {
        jamLembur = jamPerMinggu - 40
        jamNormal = 40
    } else {
        jamNormal = jamPerMinggu
        jamLembur = 0
    }

    // Menghitung total gaji bulanan
    totalGajiBulanan = (jamNormal * upahPerJam + jamLembur * 1.5 * upahPerJam) * 4

    // Tampilkan hasil
    fmt.Printf("Total gaji bulanan: Rp%.2f\n", totalGajiBulanan)
}
