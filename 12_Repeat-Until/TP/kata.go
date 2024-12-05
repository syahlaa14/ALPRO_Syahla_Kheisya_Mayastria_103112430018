package main
import "fmt"

func main() {
	var input string

	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&input)
		fmt.Println("Anda mengetik:", input)

		if input == "telkom" {
			break 
		}
	}
	fmt.Println("Program selesai.")
}
