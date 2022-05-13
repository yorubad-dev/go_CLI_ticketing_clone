package main

import (
	"fmt"
	"strings"
	"time"
	)

const conferenceTickets int = 50

var conferenceName string = "King Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	first_name string
	last_name string
	email string
	numberOfTicket uint
}

func main() {

	// greeting
	greetUser()

	for {
		// user input
		firstName, lastName, email, userTicket := userInput()

		// user input validation
		validNames, validEmail, validTickets := userInputValidation(firstName, lastName, email, userTicket)

		if validNames && validEmail && validTickets {
			// bookings
			bookingTicket(firstName, lastName, email, userTicket)

			// sending ticket simualation
			go delaySimulation(firstName, lastName, userTicket, email)
			
			// firstnames
			first_name := getFirstName()
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

func greetUser() {
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

func userInputValidation(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	validNames := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validTickets := remainingTickets > 0 && userTicket <= remainingTickets

	return validNames, validEmail, validTickets
}

func getFirstName() []string {
	first_names := []string{}
	for _, booking := range bookings {
		first_names = append(first_names, booking.first_name)
	}
	return first_names
}

func bookingTicket(firstName string, lastName string, email string, userTicket uint) {
	userData := UserData{
		first_name: firstName,
		last_name: lastName,
		email: email,
		numberOfTicket: userTicket,
	}
	// userData["FirstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTicket
	fmt.Printf("Thank you %v %v for purchasing %v tickets, confirmation will be sent to %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v ticket purchased, %v tickets remaining \n", userTicket, remainingTickets)
}

func delaySimulation(firstName string, lastName string, userTicket uint, email string)  {
	time.Sleep(5 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName )
	println("###############")
	fmt.Printf("Sending: \n %v \n to email address %v \n", ticket, email)
	println("###############")
}
