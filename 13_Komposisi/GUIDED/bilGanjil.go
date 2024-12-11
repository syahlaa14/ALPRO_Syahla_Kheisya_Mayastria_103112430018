package main
import "fmt"

func main() {
	var bil, x int
	fmt.Scan(&bil)

	for x = 1; x <= bil; x+=1 {
		if x % 2 != 0 {
			fmt.Print(x," ")
		}
	}
}