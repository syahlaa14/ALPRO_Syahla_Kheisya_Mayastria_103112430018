package main
import "fmt"

func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	if target <= 0 {
		fmt.Println("Target donasi harus berupa bilangan positif.")
		return
	}

	totalDonasi := 0
	jumlahDonatur := 0

	fmt.Println("Masukkan jumlah donasi:")
	for {
		var donasi int
		fmt.Printf("Donatur %d: ", jumlahDonatur+1)
		fmt.Scan(&donasi)

		if donasi <= 0 {
			fmt.Println("Donasi harus berupa bilangan positif.")
			continue
		}

		jumlahDonatur++
		totalDonasi += donasi

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)

		if totalDonasi >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
