package main

import ( 
	"fmt"
	"strings"
)


func main(){
	conferenceName := "Eleah's Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	bookings := []string{}
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here to attend\n", conferenceTickets)

	for {
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

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			break
		}
		remainingTickets = remainingTickets - userTickets

		//bookings array
		bookings = append(bookings, firstName + " " + lastName)
		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first booking: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
		
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are : %v\n", firstNames)
		//create if else to see if any remaining tickets left if not end loop
		if remainingTickets == 0 {
			//end program
			fmt.Printf("Our conference is booked out. Come back next year: %v\n", firstName)
			break
		}
	}

	}	
