package main

import (
	"fmt"
)

func main() {
	var jumlahOrang int

	fmt.Scan(&jumlahOrang)

	jumlahMotor := jumlahOrang / 2
	if jumlahOrang%2 != 0 {
		jumlahMotor++
	}
	
	println("Jumlah motor yang diperlukan:", jumlahMotor)
}
