package main

// First thing I did was run mod init booking-app.

// Declare imports
import (
	// Packages are collection of source files
	// for me to use with my application
	"booking-app/concurrency" // Another of my custom packages
	"booking-app/welcome"     // My own custom package
	"fmt"
	"strconv"
	"strings"
	// Possible to use WaitGroup to have concurrency
	// without an infinite for-loop.
)

// Package level constants
const currentCompany string = "Schmidt Co."
const currentDevelopers string = "Flavicorn Bros"

var userInfoBookings = make([]UserInformation, 0)

// struct: allows custom types and mixed data types
type UserInformation struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint8
}

func getCityChoice(city string) int {
	switch city {
	case "Vancouver":
		return 20
	case "Dallas", "Denver City":
		return 8
	case "Toronto":
		return 12
	case "Calgary":
		return 15
	default:
		fmt.Println("No city selected")
		return -1
	}
}

// Way to define a slice below
func getFirstNames(bookings []map[string]string) []string {
	var firstNamesSlice = []string{}

	// Can't define type here!
	for _, booking := range bookings {
		firstNamesSlice = append(firstNamesSlice, booking["firstName"]) // Like Python!
	}

	return firstNamesSlice
}

// Getting the email from a struct
func getStructEmail(userMap UserInformation) string {
	return userMap.email
}

func getUserInputs() (string, string, string, uint8, string) {
	// *Must* explicity define the type!
	var firstName string // pre-initialized variable for users name
	var lastName string
	var email string
	var userTickets uint8 // pre-initialized variable for users name
	var city string

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	if len(firstName) < 2 || len(lastName) < 2 {
		fmt.Println("Please enter at least 2 characters for your first and last name.")
		return "-1", "-1", "-1", 0, "-1"
	}

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	if !strings.Contains(email, "@") {
		fmt.Println("Please enter the correct email.")
		return "-1", "-1", "-1", 0, "-1"
	}

	fmt.Println("Please enter your city choice:")
	fmt.Scan(&city)

	var cityPrice int = getCityChoice(city)
	if cityPrice == -1 {
		fmt.Println("No such city. Please try again.")
		return "-1", "-1", "-1", 0, "-1"
	}
	fmt.Printf("You must pay %d for each ticket.", cityPrice)

	fmt.Println("Please enter your ticket amount:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets, city
}

// Entry point of the application must be declared.
// Where is the entrypoint? Where do we start the program?
func main() {

	// Go uses camelCase
	var conferenceName string = "Archibald's Clown Fest"
	var conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50
	var bookings = make([]map[string]string, 0, 50)

	// A constant in Go
	const softwareTicketLimit = 50

	welcome.WelcomeText(conferenceName, softwareTicketLimit, conferenceTickets, remainingTickets)

	// Infinite while loop
	for { // Possible to add a condition here (before the {)
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")

		firstName, lastName, email, userTickets, city := getUserInputs()
		if firstName == "-1" {
			continue
		}
		if userTickets > remainingTickets || userTickets <= 0 {
			fmt.Printf("We only have %d tickets remaining, not your amount.\n", remainingTickets)
			continue // Sends us to the next iteration of the loop (doesn't go to the code below!)
		} else {
			fmt.Println("Ticket amount accepted!")
		}

		fmt.Printf("Booking for %s location\n", city)
		remainingTickets -= userTickets

		// Map in Go
		var userData = make(map[string]string) // Cannot mix different data types!
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		var userInfo = UserInformation{
			firstName:   firstName,
			lastName:    lastName,
			email:       email,
			userTickets: userTickets,
		}

		userInfoBookings = append(userInfoBookings, userInfo)
		bookings = append(bookings, userData)

		fmt.Printf("Sold to: %v with email %v\n", getFirstNames(bookings), getStructEmail(userInfo))

		// Check that tickets still exist
		if remainingTickets == 0 {
			fmt.Println("No more tickets can be sold. Thank you!")
			fmt.Println(" ")
			fmt.Printf("Thank you for using this software ordered by %s and designed by %s", currentCompany, currentDevelopers)
			break
		}

		// Debug stuff
		fmt.Printf("Struct: %v\n", userInfoBookings)
		fmt.Printf("Map: %v\n", bookings)

		fmt.Println("************")

		// Look how easy it is to do concurrency in Go!
		go concurrency.TicketSender()

		fmt.Println("************")

	}
}
