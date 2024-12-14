package main
import "fmt"

func main() {
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 1; i <= 5; i++ {
		var warna [4]string
		fmt.Printf("Percobaan %d:\n", i)

		for j := 0; j < 4; j++ {
			fmt.Scan(&warna[j])
		}

		if warna != urutanBenar {
			berhasil = false
			break
		}
	}
	fmt.Printf("BERHASIL: %t\n", berhasil)
}
