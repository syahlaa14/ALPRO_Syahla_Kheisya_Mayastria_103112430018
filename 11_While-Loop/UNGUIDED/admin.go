package main
import "fmt"

func main() {
	var akun, pw, input string
	var total int

	akun = "Admin"
	pw = "Admin"

	fmt.Println("Masukan Akun & Password: ")
	fmt.Scan(&input, &input)

	for input != akun && input != pw {

		fmt.Println("Masukan Akun & Password dengan benar: ")
		fmt.Scan(&input, &input)
		total++

	}
	fmt.Println(total, "Percobaan gagal login")

}