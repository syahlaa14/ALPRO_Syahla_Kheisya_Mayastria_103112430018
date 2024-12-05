package main
import (
	"fmt" 
	"math")

func main() {
	var bil float64
	fmt.Scan(&bil)

	if bil <= 0 {
		fmt.Println("Masukan harus berupa bilangan desimal positif.")
		return
	}

	batas := math.Ceil(bil)

	for bil < batas {
		bil = math.Round((bil+0.1)*10) / 10 
		fmt.Printf("%.1f\n", bil)
	}
}
