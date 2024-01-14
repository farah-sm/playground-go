// This code is messy, this is because it is a demo.
// we deliberarely use most the main components of a language, these are:
	// Variables of different data types,
	// Packages & imports
	// Arrays & Slices 
	// Print Functions of different types (Print, Println, Printf)
	// User input (Scan)


// Package for the main function
package main
// Import the fmt function
import "fmt"

// Entry point of the code via main func
func main() {
	
	// Variable Initialising
	hotelName := "Reer Saeds Hotel"
	const hotelRooms = 30
	var remainingRooms = 30
	// To declare a var without a value we need to specify its data type
	var userName string
	var userRooms int
	// Slice (an empty slice)
	booking := []string{}
	// We could also declare the slice as:
		// var booking []string

	// Array (the difference between array and a slice):
		// 1. You need to explicitly specify the number of elements in the array
		// 2. The way you append a n array/ slice
	randArray := [3]string{}
	// We could also declare the array as:
		// var randArray [3]string
	
	// the use of Print vs Println is as follows:
		// "Print" doesnt print a new line whereas Println displays the contents within the parenthesis 
		// Print is ideal for user prompts such as the example below
	fmt.Print("What's your name: ")
	//Scan assumes the repsonse and initiates the variable earlier declared with the value typed by the user
	fmt.Scan(&userName)

	// Printf is ideal when you want to format numbers variables/ strings. It's use case is exampled below.
	fmt.Printf("Asalamu Alaykum %v, welcome to %v. There are %v available rooms.\n", userName, hotelName, remainingRooms)
	

	fmt.Print("How many rooms would you like to book: ")
	fmt.Scan(&userRooms)

	// App logic
	remainingRooms = remainingRooms - userRooms
	// Slice usage 
		// You can see in the Slice you use the methods "append"
			// This is the dynamic alternative to the array as the Slice 
			// will assume assign the element to the index of the Slice
	booking = append(booking, "This is a Slice: " + userName + " " + hotelName)

	// Array usage 
		// In the array you must statically input each and every one element of the array
			// explicitly specifiying its index
	randArray[0] = "This is an array: " + userName + " " + hotelName

	// Print to screen
	fmt.Printf("User: %v, booked %v rooms. There are now %v rooms remaining.\n", userName, userRooms, remainingRooms)
	// Calling the array is the same as calling a slice or a variable within the fmt.Print function
	fmt.Printf("Booking Slice: %v,\n", booking)
	fmt.Printf("Booking Array: %v,\n", randArray)

}