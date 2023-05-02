package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Print("How many tickets would you like?: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("These are all our bookings: %v\n", bookings)
			fmt.Printf("Thank you %v %v for booking %v tickets.  You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are %v tickets remaining.\n", remainingTickets)

			firstNames := []string{}

			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up.  Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}

			if !isValidEmail {
				fmt.Println("Email address is doesn't contain an '@' sign.")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}

			fmt.Println("Your input is invalid.  Try again.")
		}
	}
}
