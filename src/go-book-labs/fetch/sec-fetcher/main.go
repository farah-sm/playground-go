package main
// API: https://developer.shodan.io/api

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DNSRecord struct {
	City        string   `json:"city"`
	RegionCode  string   `json:"region_code"`
	CountryCode string   `json:"country_code"`
	CountryName string   `json:"country_name"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Hostnames   []string `json:"hostnames"`
	Domains     []string `json:"domains"`
	Org         string   `json:"org"`
	ISP         string   `json:"isp"`
	IP          string   `json:"ip"`
	Ports       []int    `json:"ports"`
	LastUpdate  string   `json:"last_update"`
	DNS         struct {
		ResolverHostname string `json:"resolver_hostname"`
		Recursive        bool   `json:"recursive"`
	} `json:"dns"`
	HttpStatus int    `json:"http_status"`
	Title      string `json:"title"`
	Favicon    struct {
		Hash     int    `json:"hash"`
		Location string `json:"location"`
	} `json:"favicon"`
	SSL struct {
		Safe bool `json:"heartbleed_safe"`
	} `json:"ssl"`
}


func main() {
	api := os.Getenv("API")
	ip := flag.String("ip", "", "To use ./expose -ip='value'")
	// alert := flag.Bool("alert", false, "To use ./expose -ip=true")
	flag.Parse()

	fmt.Printf("Scanning IPs: %v\n", *ip)
	fmt.Println("-------------------------")

	// url := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", *ip, api)

	url := fmt.Sprintf("https://api.shodan.io/shodan/ports?key=%s", api)

	//fmt.Printf("\nURL: %s\n", url)

	resp, err := http.Get(url)
	if err!=nil {
		fmt.Printf("Error with getting content from URL: %v", err)
	}

	defer resp.Body.Close()


	var dnsRecords []DNSRecord
	

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// content := []byte(body)


	err = json.Unmarshal(body, &dnsRecords)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}
	fmt.Printf("\nResult: %v", dnsRecords)
	

	//Error : 
// 	2025/04/25 14:36:45 Failed to decode JSON: json: cannot unmarshal number into Go value of type main.DNSRecord
// exit status 1

	// LOGIC
}


