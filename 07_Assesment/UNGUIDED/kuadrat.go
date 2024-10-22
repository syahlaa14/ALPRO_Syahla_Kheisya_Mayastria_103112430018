package main

import "fmt"

func main() {
    var bil int
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bil)

    for i := 1; i <= bil; i++ {
        fmt.Println(i, "=", i*i)
    }
}
