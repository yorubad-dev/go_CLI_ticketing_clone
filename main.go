package main

import "fmt"

func main() {
	conferenceName := "King Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have %v tickets and %v tickets available for the %v \n",conferenceTickets,remainingTickets,conferenceName )

}
