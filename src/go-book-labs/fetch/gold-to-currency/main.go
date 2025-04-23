package main

import (
	//"encoding/json"
	"fmt"
	// "strings"
	"io"
	"net/http"
	"os"
)

func main() {
	// sec := map[string]interface{}{}
	api := os.Getenv("API")
	url := fmt.Sprintf("https://api.metalpriceapi.com/v1/latest?api_key=%s&base=XAU&currencies=QAR,GBP", api)

	resp, err := http.Get(url)
	if err !=nil {
		panic(err)
	}

	src := resp.Body
	dst := os.Stdout

	_, err = io.Copy(dst, src)
	if err !=nil {
		panic(err)
	}

	// resp.Body.Close()

	
	


}