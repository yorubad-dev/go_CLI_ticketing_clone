package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "King Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// greeting
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets available for the %v \n", conferenceTickets, remainingTickets, conferenceName)
	for {
		// user input
		firstName, lastName, email, userTicket := userInput()

		validNames := len(firstName) >= 2 && len(lastName) >= 2
		validEmail := strings.Contains(email, "@")
		validTickets := remainingTickets > 0 && userTicket <= remainingTickets
		
		if validNames && validEmail && validTickets {
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for purchasing %v tickets, confirmation will be sent to %v \n", firstName, lastName, userTicket, email)
			remainingTickets = remainingTickets - userTicket
			fmt.Printf("%v purchased, %v tickets remaining \n", userTicket, remainingTickets)
			
			// bookings 
			first_name := []string{}
			for _, booking := range bookings {
				name := strings.Fields(booking)
				first_name = append(first_name, name[0])
			}
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