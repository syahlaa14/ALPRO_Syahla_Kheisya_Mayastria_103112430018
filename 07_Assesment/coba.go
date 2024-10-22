/*package main

import (
	"fmt"
)

func main() {
    n := 10 

    for i := 0; i < n; i += 2 {
        fmt.Println(i * (n - i - 1)) 
    }
}*/


package main

import (
	"fmt"
)

func main() {
	n := 5       // Jumlah iterasi yang diinginkan
	bil := 2     // Nilai awal

	// Perulangan sebanyak n kali
	for i := 0; i <= n; i++ {
		fmt.Println(bil) // Cetak nilai
		bil = bil * 2         // Gandakan nilai setiap iterasi
	}
}



