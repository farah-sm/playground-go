package main


import "fmt"

func main() {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	greeting := "Go conference"
	const confTickets = 50
	var remainTickets uint = 50

	
	
	fmt.Print("What is your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("What is your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("What is your email: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets would you like: ")
	fmt.Scan(&userTickets)

	remainTickets = remainTickets - uint(userTickets)

	fmt.Printf("Welcome to %v, we hope you enjoy your time here\n", greeting)
	fmt.Printf("There are %v in total. There are: %v remaining.\n", confTickets, remainTickets)


	fmt.Printf("Thank you %v %v for buying: %v tickets. You will recieve a confirmation email at: %v. There are now %v tickets remaining\n", firstName, lastName, userTickets, email, remainTickets)

}