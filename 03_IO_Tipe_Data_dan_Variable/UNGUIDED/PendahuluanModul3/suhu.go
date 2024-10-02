package main

import (
	"fmt"
)

func main() {

	var suhuFahrenheit float64
    var suhuKelvin float64

	fmt.Print("Masukkan suhu Fahrenheit: ")         // Masukkan inputan dari user 
	fmt.Scan(&suhuFahrenheit)   

	suhuKelvin = (suhuFahrenheit - 32) * (5.0/9.0) + 273.15     // Rumus konversi fahrenheit ke kelvin

	fmt.Println("Suhu", suhuFahrenheit,"°F sama dengan", suhuKelvin, "K")
}
