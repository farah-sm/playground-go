package main

import (
    "bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
			fmt.Printf("A file with Kubeconfig in its name: %s \n", path)
		} else {
			// Confirms n is a regular file and not a directory
			if info.Mode().IsRegular() {
				// fmt.Printf("%s is a regular file\n", info.Name())
				filer, err := os.Open(path)
				if err != nil {
					//
				}
				defer filer.Close()

/// -------------- OPTION 1
				// b, err := os.ReadFile(path)
				// if err != nil {
				// 	//
				// }
				// s := string(b)

				// if strings.Contains(s, "kind: config") {
				// 	fmt.Printf("File contains 'Kind: Config': %s\n", path)
				// }



/// -------------- OPTION 2

			scanner := bufio.NewScanner(filer) // NewScanner uses ScanLines func which returns each line of text, stripped of any trailing end-of-line marker.
			for scanner.Scan() {
				line := scanner.Text()
				configFile := "Kind: Config"
				if strings.Contains(line, configFile) {
					fmt.Printf("The path of a file that contains config: %s. The Line that matches: %s\n", path, line)
				}
				
			 }
			

			 }
			// 		fmt.Printf("%s\n", info.Name())
			// }
			// Open file now & assess file contents for 'kind: config'
			// filer, err := os.Open(info.Name())

			// if err != nil {
			// 	println("Error234", err)
			// }
			// defer filer.Close()

			//scanner := bufio.NewScanner(filer) // NewScanner uses ScanLines func which returns each line of text, stripped of any trailing end-of-line marker.
			//fmt.Println(*scanner)

			
		}
		return nil
})
    if err != nil {
        println("Error", err)
    }

}
