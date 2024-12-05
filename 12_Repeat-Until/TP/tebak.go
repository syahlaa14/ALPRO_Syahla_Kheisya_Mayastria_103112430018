package main
import "fmt"

func main() {
	var tebakan int

	for {
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&tebakan)

		if tebakan == 10 {
			break 
		}
		fmt.Println("Tebakan Anda salah, coba lagi.")
	}
	fmt.Println("Selamat, tebakan Anda benar!")
}
