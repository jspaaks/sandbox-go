package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName string = "Go Conference"
	const nTickets uint = 50
	var nTicketsRemaining uint = nTickets

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {

		var firstName string
		var lastName string
		var emailAddress string
		var nTicketsUser uint

		if nTicketsRemaining == 0 {
			fmt.Printf("\n%v is fully booked\n", conferenceName)
			var firstNames []string
			for _, booking := range bookings {
				firstNames = append(firstNames, strings.Split(booking, " ")[0])
			}
			fmt.Printf("These are all the bookings: %v\n\n", firstNames)
			break
		}
		fmt.Printf("%d of %d tickets remaining for %s\n", nTicketsRemaining, nTickets, conferenceName)
		fmt.Println()
		fmt.Print("    What is your first name?    : ")
		fmt.Scan(&firstName)
		fmt.Print("    What is your last name?     : ")
		fmt.Scan(&lastName)
		fmt.Print("    How many tickets you want?  : ")
		fmt.Scan(&nTicketsUser)
		fmt.Print("    What is your email address? : ")
		fmt.Scan(&emailAddress)
		fmt.Printf("\n%s %s (%s) wants to book %d tickets\n", firstName, lastName, emailAddress, nTicketsUser)

		var canBook = nTicketsRemaining >= nTicketsUser

		if canBook {
			bookings = append(bookings, firstName+" "+lastName)
			nTicketsRemaining = nTicketsRemaining - nTicketsUser
		} else {
			fmt.Println("There's not enough tickets left, try again")
		}
	}
}
