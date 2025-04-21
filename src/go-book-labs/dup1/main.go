package main 

import (
	"os"
	"fmt"
	"bufio"
)

func main() {

	count := make(map[string]int)

	// Stdin is whatever is written to the terminal via the keyboard, 
	// cound also run the command "go run main.go < input.txt"
	input := bufio.NewScanner(os.Stdin)

	// For every new input/ line, put it as a key and 
	// increment the value by 1 if the same word is repeated
	for input.Scan() {
		count[input.Text()]++
	}
	
	for line, n := range count {
		if n > 1 {
			fmt.Printf("More than 3! %v, with %v counts\n", line, n )

		}
		
	}
	fmt.Printf("\nThe final HashMap is: %v", count)
	
}