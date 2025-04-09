package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"flag"
)

func main() {

// 	var files []os.FileInfo
	fileName := flag.String("file", "", "Filename pattern we're looking for, default is kubeconfig")
	startPath := flag.String("path", ".", "Where shall we scan, default is root")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Usage: go run file.go -file <filename> [-path /start/dir]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Searching for %s starting from %s...\n", *fileName, *startPath)

	err := filepath.Walk(*startPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), *fileName) {
			fmt.Printf("%s \n", path)
		}
		return nil
})
    if err != nil {
        println("Error", err)
    }

}
