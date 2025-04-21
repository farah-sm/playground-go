package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string

	strings.Join(os.Args[1:], " ")

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[1]
		sep = " "

	} 
	fmt.Printf("%s", s)
}