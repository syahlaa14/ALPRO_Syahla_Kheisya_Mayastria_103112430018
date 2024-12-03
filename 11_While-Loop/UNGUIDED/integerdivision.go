package main
import "fmt"

func main() {
	var x, y, jml int

	fmt.Print("Masukan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukan bilangan y: ")
	fmt.Scan(&y)

	for x >= y {
		x = x - y
		jml++
	}

	fmt.Print(jml)

}