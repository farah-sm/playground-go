package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get the API key from environment variable
	api := os.Getenv("API")

	// Parse command line arguments
	ip := flag.String("ip", "", "To use ./expose -ip='value'")
	flag.Parse()

	// If an IP is provided, display it
	fmt.Printf("Scanning IPs: %v\n", *ip)
	fmt.Println("-------------------------")

	// Construct the URL for the Shodan API (using the API key and the IP address)
	url := fmt.Sprintf("https://api.shodan.io/shodan/ports?key=%s", api)

	// Make an HTTP GET request to the Shodan API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error with getting content from URL: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the raw JSON response to inspect its structure
	fmt.Println("Raw JSON response:")
	fmt.Println(string(body))

	// Create a slice to hold the ports
	var ports []int

	// Unmarshal the JSON response into the ports slice
	err = json.Unmarshal(body, &ports)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	// Print the ports to verify the response
	fmt.Println("Open ports:")
	for _, port := range ports {
		fmt.Println(port)
	}

	// Additional logic can go here
}
