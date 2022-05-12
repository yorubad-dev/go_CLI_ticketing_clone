package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

func main() {
	conferenceName := "King Conference"
	var remainingTickets uint = 50
	bookings := []string{}

	// greeting
	greetUser(conferenceName, remainingTickets)

	for {
		// user input
		firstName, lastName, email, userTicket := userInput()

		// user input validation
		validNames, validEmail, validTickets := userInputValidation(firstName, lastName, email, remainingTickets, userTicket)

		if validNames && validEmail && validTickets {

			fmt.Printf("Thank you %v %v for purchasing %v tickets, confirmation will be sent to %v \n", firstName, lastName, userTicket, email)
			fmt.Printf("%v purchased, %v tickets remaining \n", userTicket, remainingTickets)
			remainingTickets = remainingTickets - userTicket

			// bookings
			bookings = append(bookings, firstName+" "+lastName)
			first_name := getFirstName(bookings)
			fmt.Printf("These are all our booking: %v\n", bookings)
			fmt.Printf("These are all the first name of the bookings: %v\n", first_name)

			if remainingTickets == 0 {
				fmt.Printf("%v has been booked out\n", conferenceName)
				break
			}
		} else {
			if !validNames {
				fmt.Println("Either first name or last name are too short")
			}
			if !validEmail {
				fmt.Println("Email does not contain @... it is invalid")
			}
			if !validTickets {
				println("Ticket entered is invalid")
			}
		}
	}
}

func greetUser(conferenceName string, remainingTickets uint) {
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets available for the %v \n", conferenceTickets, remainingTickets, conferenceName)
	// return conferenceName, remainingTickets
}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email Address: ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func userInputValidation(firstName string, lastName string, email string, remainingTickets uint, userTicket uint) (bool, bool, bool) {
	validNames := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validTickets := remainingTickets > 0 && userTicket <= remainingTickets

	return validNames, validEmail, validTickets
}

func getFirstName(bookings []string) []string{
	first_name := []string{}
	for _, booking := range bookings {
		name := strings.Fields(booking)
		first_name = append(first_name, name[0])
	}
	return first_name
}
