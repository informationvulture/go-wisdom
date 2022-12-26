package welcome

import (
	"fmt"
)

// Capital letter in front to export it.
func WelcomeText(conferenceName string, softwareTicketLimit uint8, conferenceTickets uint8, remainingTickets uint8) {
	fmt.Printf("Welcome to the %s booking system! ", conferenceName)
	fmt.Printf("As of version 0.0.3, we have a limit of %d tickets you can sell.\n", softwareTicketLimit)
	fmt.Println("Please continue to get your tickets.")
	fmt.Println(" ")
	fmt.Printf("There are %d tickets remaining, with %d tickets originally.\n", conferenceTickets, remainingTickets)
}
