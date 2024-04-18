package main

import "fmt"

func main(){
	conferenceName := "Eleah's Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here to attend\n", conferenceTickets)

	var userName string
	var userTickets int
	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
}
