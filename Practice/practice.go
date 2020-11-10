package main

import (
	"fmt"
	"time"
)

func printMe(from string) string {
	return from
}

func main() {
	go printMe("Golang")
	fmt.Print()
	printMe("async")
}
