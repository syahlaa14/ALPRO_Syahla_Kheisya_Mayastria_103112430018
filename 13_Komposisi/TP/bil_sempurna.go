package main
import "fmt"

func main() {
	var angka, jumlah int
	fmt.Scan(&angka)

	jumlah = 0
	for i := 1; i < angka; i++ {
		if angka%i == 0 {
			jumlah += i
		}
	}

	if jumlah == angka {
		fmt.Printf("%d adalah bilangan sempurna\n", angka)
	} else {
		fmt.Printf("%d bukan bilangan sempurna\n", angka)
	}
}
