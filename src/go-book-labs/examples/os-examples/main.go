package main 

import (
	"os"
	"fmt"
	
)

func main() {

	// os.Args[] is an array of parameter values, 'go run main.go Saed' 
	// 'Saed' is index 1 of the array
	name := os.Args[1]

	fmt.Printf("Hello %s\"n, name)
}