package main

import (
	"fmt"
	//"strings"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		//handle the err
		if err!= nil{
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		//handle the err
		if err!= nil{
			fmt.Printf("error: %v", err)
		}
		

		fmt.Printf("%v", b)

		



	}

}