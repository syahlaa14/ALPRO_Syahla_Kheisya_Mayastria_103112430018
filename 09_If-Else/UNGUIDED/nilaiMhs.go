package main

import (
	"fmt"
)

func main() {
	var nilai int
	fmt.Scan(&nilai)

	if nilai > 90 {
		fmt.Println("A")
	} else if nilai >= 80 && nilai <= 90 {
		fmt.Println("AB")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
