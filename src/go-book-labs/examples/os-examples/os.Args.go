package main 

import (
	"os"
	"fmt"
	
)

func main() {

	name := os.Args[1]

	fmt.Printf("Hello %s\n", name)
}