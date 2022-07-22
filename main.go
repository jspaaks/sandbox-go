package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Booking struct {
	firstName    string
	lastName     string
	emailAddress string
	nTickets     uint
}

const conferenceName string = "Go Conference"

var bookings = make([]Booking, 0)

func main() {

	const nTickets uint = 50
	var nTicketsRemaining uint = nTickets
	var emailConfirmationWaitGroup = sync.WaitGroup{}

	greetUsers()

	for {
		fmt.Printf("\n%d of %d tickets remaining for %s\n\n", nTicketsRemaining, nTickets, conferenceName)

		var booking Booking = getUserInput()

		var canBook = validateUserInput(booking, nTicketsRemaining)
		if canBook {
			nTicketsRemaining = updateBookings(booking, nTicketsRemaining)
			emailConfirmationWaitGroup.Add(1)
			go func(booking Booking) {
				defer emailConfirmationWaitGroup.Done()
				sendEmail(booking)
			}(booking)
		} else {
			fmt.Printf("    Please try again.\n\n")
		}
		if nTicketsRemaining == 0 {
			printFullyBooked()
			break
		}
	}

	emailConfirmationWaitGroup.Wait()

}

func greetUsers() {
	fmt.Printf("\n")
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("Get your tickets here to attend\n")
	fmt.Printf("\n")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() Booking {
	var firstName string
	var lastName string
	var emailAddress string
	var nTicketsUser uint

	fmt.Print("    What is your first name?    : ")
	fmt.Scan(&firstName)
	fmt.Print("    What is your last name?     : ")
	fmt.Scan(&lastName)
	fmt.Print("    How many tickets you want?  : ")
	fmt.Scan(&nTicketsUser)
	fmt.Print("    What is your email address? : ")
	fmt.Scan(&emailAddress)
	fmt.Printf("\n    %s %s (%s) wants to book %d tickets.\n", firstName, lastName, emailAddress, nTicketsUser)

	var booking = Booking{
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: emailAddress,
		nTickets:     nTicketsUser,
	}

	return booking
}

func updateBookings(booking Booking, nTicketsRemaining uint) uint {
	bookings = append(bookings, booking)
	nTicketsRemaining = nTicketsRemaining - booking.nTickets
	fmt.Printf("    %d tickets booked.\n", booking.nTickets)
	return nTicketsRemaining
}

func printFullyBooked() {
	fmt.Printf("\n%v is now fully booked\n", conferenceName)
	fmt.Printf("These are all the bookings: %v\n\n", getFirstNames())
}

func validateUserInput(booking Booking, nTicketsRemaining uint) bool {

	var nameIsValid bool = len(booking.firstName) >= 2 && len(booking.lastName) >= 2
	var emailAddressIsValid bool = strings.Contains(booking.emailAddress, "@")
	var nTicketsUserIsValid bool = booking.nTickets > 0
	var enoughTicketsAvailable bool = nTicketsRemaining >= booking.nTickets
	var canBook = nameIsValid && emailAddressIsValid && nTicketsUserIsValid && enoughTicketsAvailable

	if !canBook {
		fmt.Printf("\n    Detected a problem with your booking:\n")
	}
	if !nameIsValid {
		fmt.Println("    - First name or Last name too short")
	}
	if !emailAddressIsValid {
		fmt.Println("    - Invalid email address")
	}
	if !nTicketsUserIsValid {
		fmt.Println("    - Invalid number of tickets")
	}
	if !enoughTicketsAvailable {
		fmt.Println("    - There's not enough tickets left")
	}

	return canBook
}

func sendEmail(booking Booking) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", booking.nTickets, booking.firstName, booking.lastName)
	fmt.Printf("\n    Sent %v to email address %v.\n\n", ticket, booking.emailAddress)
}
