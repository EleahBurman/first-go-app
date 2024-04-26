package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Eleah's Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	bookings := []string{}

	greetUsers()
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here to attend\n", conferenceTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			// Check if all tickets are booked
			if remainingTickets == 0 {
				//end program
				fmt.Printf("Our conference is booked out. Come back next year.\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets your entered is invalid")
			}
		}
	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	//execute code for booking New York conference
	// case "Singapore", "Hong Kong":
	// 	//execute code for booking Singapore & Hong Honk conference
	// case "London", "Berlin":
	// 	//execute code for booking London & Berlin conference
	// case "Mexico City":
	// 	//execute code for booking Mexico City conference
	// default:
	// 	fmt.Print("No valid city selected")
	// }
}


func greetUsers(){
	fmt.Println("Welcome to our conference")
}