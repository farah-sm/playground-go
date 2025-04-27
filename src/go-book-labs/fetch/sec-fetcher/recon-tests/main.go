package main
// API: https://developer.shodan.io/api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//"flag"
)


func main() {
	api := os.Getenv("API")
	var responz ApiResponse

	// url := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", *ip, api)
	

	url := fmt.Sprintf("https://api.shodan.io/shodan/host/search?key=%s&query=http.title:grafana", api)

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
	err = json.Unmarshal(body, &responz)
	if err != nil {
			log.Fatalf("Failed to decode JSON: %v", err)
	}

	// fmt.Println("-------------------------")
	// fmt.Printf("Found: %v\n", responz)
	// fmt.Println("-------------------------")

	count := make(map[string]int64)

	for i:=0; i < len(responz.Matches); i++ {
		// a:=i+1
		// fmt.Printf("Country number: %d is %v. City %v.\n", a, responz.Matches[i].Location.CountryName, responz.Matches[i].Location.City)

		for _ = range responz.Matches[i].Location.CountryName {
			count[responz.Matches[i].Location.CountryName]++
		}
	
	}
	fmt.Printf("%v\n", count)
	


}