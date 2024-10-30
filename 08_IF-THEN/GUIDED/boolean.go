package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	if angka %2 == 0 {
		fmt.Println("true")
	}
	if angka %2 != 0 {
		fmt.Println("false")
	}
}
