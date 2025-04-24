package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ApiResponse struct {
	Base string `json:"base"`
	Timestamp int64 `json:"timestamp"`
	Rates map[string]float64 `json:"rates"`
}

func main() {

	//counts := make(map[string]int)
	currency := flag.String("currency", "XAU", "Currency, '-currency=XAU,EUR'")
	base := flag.String("base", "GBP", "-base=GBP")
	api := os.Getenv("API")
	flag.Parse()
	url := fmt.Sprintf("https://api.metalpriceapi.com/v1/latest?api_key=%s&base=%s&currencies=%s", api, *base,*currency)

	// fmt.Printf("%s", url)
	resp, err := http.Get(url)
	if err !=nil {
		fmt.Printf("Error: %s", err.Error())
	}

    var result ApiResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Fatalf("Failed to decode JSON: %v", err)
    }

    fmt.Printf("Base: %s\n", result.Base)
    fmt.Printf("Timestamp: %d\n", result.Timestamp)
    fmt.Println("Rates:")
    for curr, val := range result.Rates {
        fmt.Printf("  %s: %.10f\n", curr, val)
    }
	


}