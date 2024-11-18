package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string

	fmt.Scan(&huruf)

	huruf = strings.ToUpper(huruf)

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan")
	}
}
