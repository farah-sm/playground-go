package main

import (
	"fmt"
	"net/http"
	"strings"
	"flag"
	
)

func main() {
	// Flag for the username in question
	user := flag.String("user", "me", "Username in question")
	flag.Parse()
	prefix := "https://"
	fullPath := make(map[string]int)

	platform := map[string]int{
		"instagram.com":1,
		"facebook.com":2,
		"linkedin.com/in":3,
		"tiktok.com":4,
	}

	for p, _ := range platform {
		holder := prefix + p + "/" + *user

		fullPath[holder]++
	}


	// HTTP GET
	for url, _ := range fullPath {
		resp, err := http.Get(url)
		// handle the err
		if err!= nil{
			fmt.Printf("error: %v", err)
		}

		status := resp.Status

		if strings.Contains(status, "200") {
			fmt.Printf("Success! Code: %s ", status)
			fmt.Printf("HIT: %s is a valid URL with %s\n", *user, url)

		} else {
			fmt.Printf("Unsuccessful! Code: %s ", status)
			fmt.Printf("NOT A HIT: %s is an invalid URL with %s\n", *user, url)
		}

	}

}