package main

import "fmt"

func main() {
	conferenceName := "King Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// greeting
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets available for the %v \n",conferenceTickets,remainingTickets,conferenceName )

	// user input
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

	fmt.Printf("Thank you %v %v for purchasing %v tickets, confirmation will be sent to %v \n", firstName, lastName,userTicket, email)
	remainingTickets = remainingTickets - userTicket
	fmt.Printf("%v purchaseed, %v tickets remaining \n",userTicket, remainingTickets)
}
