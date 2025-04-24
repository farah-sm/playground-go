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
	//fullPath := make(map[string]int)

	fmt.Printf("Searching the web, hang tight...\n\n")

	platform := map[string]string{
		"Instagram":"https://www.instagram.com/",
		"Facebook":"https://www.facebook.com/",
		"YouTube":"https://www.youtube.com/user/",
		"Reddit":"https://www.reddit.com/user/",
		"GitHub":"https://www.github.com/",
		"Twitch":"https://www.twitch.tv/",
		"Pinterest":"https://www.pinterest.com/",
		"TikTok":"https://www.tiktok.com/@",
		"Flickr":"https://www.flickr.com/photos/",
	}
	var url string

	for platform, link := range platform { {
		url = link + *user
		//fmt.Printf("\nURL IS: %s\n", url)
		// url := fullPath[holder]++
		// fmt.Printf("%s\n", holder)
		resp, err := http.Get(url)
		// handle the err
		if err!= nil{
			fmt.Printf("error: %v", err)
		}

		status := resp.Status

		if strings.Contains(status, "200") {
			fmt.Printf("❌ %s: @%s is unavailable. URL: %s. Status: %s\n", platform, *user, url, status)

		} else if strings.Contains(status, "404") {
			fmt.Printf("✅ %s: @%s is available. URL %s. Status: %s\n", platform, *user, url, status)
		} else {
			fmt.Printf("🤔 %s @%s is unclear if available. URL %s. Status: %s\n", platform, *user, url, status)
		}
		
	}
}


	// HTTP GET


	

}