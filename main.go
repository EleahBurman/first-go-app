package main

import "fmt"

func main(){
	var conferenceName = "Eleah's Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here to attend", conferenceTickets)

	var userName string
	//ask user for their name
	var userTickets int
	userName = "Tom"
	userTickets = 2
	fmt.Println(userName)
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
}
