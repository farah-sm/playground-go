package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What's your name: ")
	scanner.Scan()
	name := scanner.Text()

	writer := bufio.NewWriter(os.Stdout)

	fmt.Fprintf(writer, "Hello ", name)
	writer.Flush()


}