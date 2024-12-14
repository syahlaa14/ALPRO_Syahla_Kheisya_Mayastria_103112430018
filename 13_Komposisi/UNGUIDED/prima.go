package main
import "fmt"

type CekPrima struct {
	n int
}

func (cp CekPrima) IsPrima() bool {
	if cp.n <= 1 {
		return false
	}
	for i := 2; i*i <= cp.n; i++ {
		if cp.n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	
	cek := CekPrima{n: n}
	if cek.IsPrima() {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
