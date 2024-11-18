package main

import (
	"fmt"
	"strings"
)

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Scan(&umur)

	fmt.Scan(&kewarganegaraan)

	kewarganegaraan = strings.ToUpper(kewarganegaraan)

	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else if umur < 17 && kewarganegaraan != "WNI" {
		fmt.Println("Umur kurang dari 17 tahun dan bukan WNI")
	} else if umur < 17 {
		fmt.Println("Umur Anda kurang dari 17 tahun")
	} else {
		fmt.Println("Bukan WNI")
	}
}
