package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= int(remainTickets)

	return isValidName, isValidEmail, isValidTickets
}