package main
import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)

	fmt.Printf("Bilangan prima dari 1 - %d adalah:\n", bil)
	for i := 2; i <= bil; i++ {
		prima := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prima = false
				break
			}
		}
		if prima {
			fmt.Printf("%d ", i)
		}
	}
}
