package main

import (
	"fmt"
	"strings"
)
var greeting = "Go conference"
const confTickets = 50
var remainTickets uint = 50
var bookings = []string{}

func main() {


	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput() 

		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, int(userTickets))

		if isValidName && isValidEmail && isValidTickets {

			// book ticket
			bookTicket(userTickets, firstName, lastName, email)

			if remainTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Please try again next year.")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("Incorrect Name")
			}
			if !isValidTickets {
				fmt.Println("Incorrect Ticket")
			}
			if !isValidEmail {
				fmt.Println("Incorrect Email")
			}
			// skips all the follwong code, but goes to the next iteration of the loop.
			// continue
		}

	}
}

func greetUser() {
	fmt.Printf("Welcome to %v, we hope you enjoy your time here\n", greeting)
	fmt.Printf("There are %v in total. There are: %v remaining.\n", confTickets, remainTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	// this is a for each loop, taking both the index and the value of the slice
	for _, booking := range bookings {
		// we split the booking variable (full name of the user), and store it in a 'name' slice
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= int(remainTickets)

	return isValidName, isValidEmail, isValidTickets
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("What is your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("What is your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("What is your email: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets would you like: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string ) {

	remainTickets = uint(remainTickets) - uint(userTickets)

	bookings = append(bookings, firstName+" "+lastName)

	// we take the return value of the function 'getFirstNames' and store it in a variable
	firstNames := getFirstNames()

	fmt.Printf("The first names of our users are: %v\n", firstNames)
	fmt.Printf("Thank you %v %v for buying: %v tickets. You will recieve a confirmation email at: %v. There are now %v tickets remaining\n", firstName, lastName, userTickets, email, remainTickets)

}