package main

import "fmt"

func main() {
    var bil int     // deklarasi variabel
    fmt.Scan(&bil)

    for i := 1; i <= bil; i++ {         //menghitung bilangan kuadrat dari 1 sampai inputan user
        fmt.Println("Hasil kuadrat dari", i, "=", i*i)
    }
}
