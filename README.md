# Basic Concepts in Go:

## Print Functions:
1. **`println` vs. `print` vs. `printf`:**
   - `println`: Outputs a line of values, automatically adding a newline character.
   - `print`: Outputs values without adding a newline.
   - `printf`: Formats and prints values based on a format specifier (like in C's printf).

## Variables:
2. **Variable Declaration:**
   - `var variableName dataType`: Declares a variable with a specified data type.
   - Example: `var age int` declares an integer variable named `age`.

## Arrays and Slices:
3. **Array vs. Slice:**
   - **Array:** Fixed-size collection of elements.
     - Declaration: `var arr [3]int` creates an array of three integers.
   - **Slice:** Dynamic, flexible view of an array.
     - Declaration: `var mySlice []int` creates an empty slice.

4. **Appending to Slice vs. Array:**
   - **Slice:**
     - `mySlice = append(mySlice, 42)`: Adds an element to the end of the slice.
   - **Array:**
     - You can't directly append to an array; its size is fixed.

## Packages:
5. **Package in Go:**
   - A package is a way to organize and reuse code.
   - `main` package is special and needed for executable programs.
   - `package main` at the top of the file indicates it's an executable program.

## Imports:
6. **Importing Packages:**
   - `import "fmt"`: Imports the "fmt" package for formatted I/O.
   - You need to import a package to use functions or features from that package.

### Examples:

```go
// Variable Declaration
var age int

// Array Declaration
var myArray [3]int

// Slice Declaration
var mySlice []int

// Append to Slice
mySlice = append(mySlice, 42)

// Print Functions
fmt.Println("Hello, Go!") // println
fmt.Print("No newline")   // print
fmt.Printf("Formatted: %d", 42) // printf

// Package and Import
package main

import "fmt"

func main() {
    // Your main program logic goes here
}

