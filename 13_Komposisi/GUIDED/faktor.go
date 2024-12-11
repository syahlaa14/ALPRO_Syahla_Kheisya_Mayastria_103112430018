package main  
import "fmt"  

func main() {  
    var bilangan, x int 
	fmt.Scan(&bilangan) 

    for x = 1; x <= bilangan; x += 1 { 
        if bilangan % x == 0 { 
            fmt.Print(x," ") 
        } 
    } 
}