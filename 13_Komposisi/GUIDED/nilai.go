package main
import "fmt"

func main() {
	var bil1, bil2, bil3, max, min int
	fmt.Scan(&bil1, &bil2, &bil3)

	if bil1 > bil2 {
		max = bil1
		min = bil2
	} else {
		max = bil2
		min = bil1
	}

	if max < bil3 {
		max = bil3
	}

	if min > bil3 {
		min = bil3
	}
	fmt.Println("Terbesar:", max)
	fmt.Println("Terkecil:", min)
}
