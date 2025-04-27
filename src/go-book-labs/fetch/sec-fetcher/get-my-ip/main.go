package main
// API: https://developer.shodan.io/api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)
	var myip string 

func main() {
	api := os.Getenv("API")

	// url := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", *ip, api)

	url := fmt.Sprintf("https://api.shodan.io/tools/myip?key=%s", api)

	//fmt.Printf("\nURL: %s\n", url)

	resp, err := http.Get(url)
	if err!=nil {
		fmt.Printf("Error with getting content from URL: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
			log.Fatalf("Error reading response body: %v", err)
	}

	// Unmarshal the JSON response into the ports slice
	err = json.Unmarshal(body, &myip)
	if err != nil {
			log.Fatalf("Failed to decode JSON: %v", err)
	}

	fmt.Println("-------------------------")
	fmt.Printf("Your public IP: %v\n", myip)
	fmt.Println("-------------------------")
}