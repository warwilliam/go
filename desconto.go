package main

import "fmt"




func finalPrice  (name string,price float64,disccount float64){
	
	 finalPrice := price * (1 - disccount)

	 fmt.Print("o preco final do "+ name+" e de R$: ")
	 fmt.Println(finalPrice)
	}






