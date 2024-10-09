package main

import (
	"fmt"
)

func main() {
	var angka, digit1, digit2, digit3 int			//deklarasi variabel
	fmt.Print("Masukkan bilangan 3 digit: ")		// Input bilangan dari user
	fmt.Scan(&angka)

	digit1 = angka / 100            
	digit2 = (angka / 10) % 10       
	digit3 = angka % 10               

	fmt.Println(digit1 <= digit2 && digit2 <= digit3)
}
