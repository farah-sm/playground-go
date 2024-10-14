package main


import "fmt"


func main() {
	p, j := 0, 55
	// the value of i is the memory location of p
	i := &p
	// we change the value of p by reference, using the i variable
	*i = 22
	// the value of p is now 22 (after line 13), j is still 55
	// and the value of the i is still the hexadecimal memory location of the variable p
	fmt.Print(p, j, i)
}
