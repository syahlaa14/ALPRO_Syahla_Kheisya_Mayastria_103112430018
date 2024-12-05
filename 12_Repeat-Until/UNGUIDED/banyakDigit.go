package main
import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	if angka <= 0 {
		fmt.Println("Masukan bilangan bulat positif: ")
		return
	}

	hitung := 0
	temp := angka

	for {
		hitung++
		temp /= 10
		if temp == 0 {
			break 
		}
	}

	fmt.Println(hitung)
}
