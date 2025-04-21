package main
 
import (
	"fmt"
    "strings"
	// "os"
	// "io/utils"
)

func main() {
	name := []string{"saed", "ali", "yusuf"}

	names := strings.Join(name, " ")
	fmt.Println(names)
}