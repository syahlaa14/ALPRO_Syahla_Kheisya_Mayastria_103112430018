package main

import (
	"fmt"
)

func main() {

	var sisi int
	var volume int

	fmt.Println("Masukan sisi kubus: ") 
	fmt.Scan(&sisi)

	volume = sisi * sisi * sisi

	fmt.Println("volume kubus adalah: ", volume)
}
