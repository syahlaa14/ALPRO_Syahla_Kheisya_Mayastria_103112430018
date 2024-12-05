package main
import "fmt"

func main() {
	var harga, total int
	fmt.Println("Masukkan harga barang (Ketik 0 untuk selesai):")

	for {
		fmt.Print("Masukkan harga barang: ")
		fmt.Scan(&harga)

		if harga == 0 {
			break
		}
		total += harga
	}
	fmt.Println("Total harga belanja Anda adalah", total)
}
