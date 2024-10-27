package main 

import ( "fmt" ) 

func main() { 
	var bilangan int 
	fmt.Print("Masukkan bilangan tiga digit: ") 
	fmt.Scan(&bilangan) 

	digitPertama := bilangan / 100 
	digitKedua := (bilangan / 10) % 10 

	fmt.Println(digitPertama > digitKedua && digitKedua > digitKetiga) 
}