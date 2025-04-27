package main
// API: https://developer.shodan.io/api

import (
	//"encoding/json"
	"fmt"
	"io"
	"time"
	"net/http"
	"os"
)
	// var myip string 

	

	func fetch(url string, ch chan <- string) {
		start := time.Now()

		resp, err := http.Get(url)
		if err!=nil {
			ch <- fmt.Sprintf("Error getting URL: %v", err)
			return		
		}
	
		_, err = io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
				ch <- fmt.Sprintf("Error reading response body: %v", err)
				return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2fs %s ", secs, url)

	}

func main() {


	// api := os.Getenv("API")

	start:=time.Now()
	ch := make(chan string)

	// url := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", *ip, api)
	// url := fmt.Sprintf("https://api.shodan.io/shodan/host/185.230.63.171?key=%s/", api)

	//fmt.Printf("\nURL: %s\n", url)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
	fmt.Println(<- ch)
	}
	elapsed := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", elapsed)
	




	// Unmarshal the JSON response into the ports slice
	// err = json.Unmarshal(body, &myip)
	// if err != nil {
	// 		log.Fatalf("Failed to decode JSON: %v", err)
	// }

	// fmt.Println("-------------------------")
	// fmt.Printf("Your public IP: %v\n", myip)
	// fmt.Println("-------------------------")


}
