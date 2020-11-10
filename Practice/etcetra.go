package main

import (
	"fmt"
)

type Biscuit struct{
	Brand 		string
	ProductName	string
	Price		int
}

func main() {
	// var biscuit Biscuit

	biscuit:= Biscuit{"Cremica", "Jeera Lite", 5,}
	fmt.Print(biscuit.Price)

	
}
