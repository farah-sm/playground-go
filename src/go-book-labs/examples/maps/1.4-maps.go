package main

import (
	"fmt"

)


func main() {
	
	counts := make(map[string]int)

	counts["hello"] = 1
	counts["Saed"] = 7

	//fmt.Println(counts["hello"])

	for i, m := range counts {
		fmt.Println(i, m)
	}
}