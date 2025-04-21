package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	//"strings"
)

func main() {
	for _, url := range os.Args[1:] {
	//  prefix := strings.HasPrefix(url, "https://")
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("err: %s", err.Error())
		}
		src := resp.Body
		dst := os.Stdout
		b, err := io.Copy(dst, src)
		status := resp.Status
		
	//io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n\n\n Status: ", b, status)
	}
}