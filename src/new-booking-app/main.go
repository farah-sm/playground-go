package main


import "fmt"


func main() {
	p, j := 0, 55

	// the value of i is the memory location of p
	i := &p
	// we change the value of p by reference, using the i variable
	*i = 22



	fmt.Print(p, j, i)
}
