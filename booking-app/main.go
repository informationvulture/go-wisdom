package main

// First thing I did was run mod init booking-app.

// Declare imports
import (
	// Packages are collection of source files
	// for me to use with my application
	"fmt"
	"strings"
)

// Entry point of the application must be declared.
// Where is the entrypoint? Where do we start the program?
func main() {

	// Go uses camelCase
	var conferenceName string = "Archibald's Clown Fest"
	var conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50
	var bookings = []string{} // Slice since we didn't specify the length

	// A constant in Go
	const softwareTicketLimit = 50

	fmt.Printf("Welcome to the %s booking system! ", conferenceName)
	fmt.Printf("As of version 0.0.2, we have a limit of %d tickets you can sell.\n", softwareTicketLimit)
	fmt.Println("Please continue to get your tickets.")
	fmt.Println(" ")
	fmt.Printf("There are %d tickets remaining, with %d tickets originally.\n", conferenceTickets, remainingTickets)

	// Infinite while loop
	for { // Possible to add a condition here (before the {)
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		// *Must* explicity define the type!
		var firstName string // pre-initialized variable for users name
		var lastName string
		var email string
		var userTickets uint8 // pre-initialized variable for users name

		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		if len(firstName) < 2 || len(lastName) < 2 {
			fmt.Println("Please enter at least 2 characters for your first and last name.")
			continue
		}

		fmt.Println("Please enter your email:")
		fmt.Scan(&email)

		if !strings.Contains(email, "@") {
			fmt.Println("Please enter the correct email.")
			continue
		}

		fmt.Println("Please enter your ticket amount:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets || userTickets <= 0 {
			fmt.Printf("We only have %d tickets remaining.\n", remainingTickets)
			continue // Sends us to the next iteration of the loop (doesn't go to the code below!)
		} else {
			fmt.Println("Ticket amount accepted!")
		}
		remainingTickets -= userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("User %s %s wants %d tickets\n", firstName, lastName, userTickets)
		fmt.Printf("THANK YOU! Confirmation sent to %s for %d tickets\n", email, userTickets)
		fmt.Printf("There are now %d tickets for %s\n", remainingTickets, conferenceName)

		var firstNamesSlice = []string{}

		// Can't define type here!
		for _, booking := range bookings {
			firstNamesSlice = append(firstNamesSlice, strings.Fields(booking)[0])
		}
		fmt.Printf("%v\n", firstNamesSlice)

		// Check that tickets still exist
		if remainingTickets == 0 {
			fmt.Println("No more tickets can be sold. Thank you!")
			break
		}

		// Stopped at 1:43:48
	}
}
