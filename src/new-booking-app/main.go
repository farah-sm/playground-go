package main

import (
	"fmt"
	"time"
	// "sync"

)
var greeting = "Go conference"
const confTickets = 50
var remainTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}



func main() {


	greetUser()
	
	//  
	// var wg =: sync.WaitGroup{}

	for {

		firstName, lastName, email, userTickets := getUserInput() 

		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, int(userTickets), remainTickets)

		if isValidName && isValidEmail && isValidTickets {

// -->		// book ticket
			bookTicket(userTickets, firstName, lastName, email) 
// -->		// the below function calls a number of threads this main func should wait for, 
// --> 		// increasing the counter of threads the app should wait for
			// wg.Add(1)
			// concurrency key word "go"
			go sendEmail(userTickets, firstName, lastName, email)

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
	// Waits for all the threads that were added before they are done
	// wg.Wait()
	

}

func greetUser() {
	fmt.Printf("Welcome to %v, we hope you enjoy your time here\n", greeting)
	fmt.Printf("There are %v in total. There are: %v remaining.\n", confTickets, remainTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	// this is a for each loop, taking both the index and the value of the slice
	for _, booking := range bookings {
		// we take the value of the key, "firstName", and store it in a 'firstName' slice
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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
	
	// create a map for a user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}


	bookings = append(bookings, userData)
	
	fmt.Printf("List of booking is: %v\n", bookings)

	// we take the return value of the function 'getFirstNames' and store it in a variable
	firstNames := getFirstNames()

	fmt.Printf("The first names of our users are: %v\n", firstNames)
	fmt.Printf("Thank you %v %v for buying: %v tickets. You will recieve a confirmation email at: %v. There are now %v tickets remaining\n", firstName, lastName, userTickets, email, remainTickets)

}

func sendEmail(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n%v to email address:\n%v\n", ticket, email)
	fmt.Println("##############")

	// This decrements the waiting function by 1
	// wg.Done(0)
}