package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"flag"
	"bufio"
)

func main() {

	counts := make(map[string]int)
	currency := flag.String("currency", "GBP", "Currency, '-currency=USD,GBP'")
	api := os.Getenv("API")
	flag.Parse()
	url := fmt.Sprintf("https://api.metalpriceapi.com/v1/latest?api_key=%s&base=XAU&currencies=%s", api, *currency)

	fmt.Printf("%s", url)
	resp, err := http.Get(url)
	if err !=nil {
		panic(err)
	}

	file, _ := os.Open("file.txt")


	

	src := resp.Body
	// dst := file

	io.Copy(file, src)
	// if err !=nil {
	// 	panic(err)
	// }
	input := bufio.NewScanner(file)
	defer file.Close()


	for input.Scan() {
		counts[input.Text()]++
	}




	for line, _ := range counts {
		fmt.Printf("\n%s\n", line)
	}

	resp.Body.Close()

	
	


}