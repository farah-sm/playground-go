package main

package main


import (
	// "k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
	"fmt"
	"os"
	"strings"
	// "io/fs"
	"path/filepath"
	"flag"
)



func main() {

	var files []os.FileInfo
	fileName := flag.String("file", "", "Filename pattern we're looking for, default is kubeconfig")
	startPath := flag.String("path", "/", "Where shall we scan, default is root")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Usage: go run file.go -file <filename> [-path /start/dir]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Searching for %s starting from %s...\n", *fileName, *startPath)

	//fmt.Printf("%v", *fileName)
	err := filepath.Walk(*startPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), *fileName) {
			files = append(files, info)
		}
		return nil
})
    if err != nil {
        println("Error", err)
    } else {
        for _, f := range files {
            println(f.Name())
            // This is where we'd like to open the file
        }


}

}


}

}
