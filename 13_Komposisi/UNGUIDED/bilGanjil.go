package main
import "fmt"


type HitungGanjil struct {
	n int
}

func (h HitungGanjil) Hitung() int {
	count := 0
	for i := 1; i <= h.n; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Masukan harus bilangan bulat positif.")
		return
	}

	penghitung := HitungGanjil{n: n}
	hasil := penghitung.Hitung()

	fmt.Printf("Terdapat %d bilangan ganjil\n", hasil)
}
