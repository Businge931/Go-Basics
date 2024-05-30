package main

import (
	"fmt"

	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
} 

func main (){
	userfirstName :=getUserData("Please enter your first name: ")
	userlastName :=getUserData("Please enter your last name: ")
	userbirthdate :=getUserData("Please enter your bithdate (MM/DD/YYYY): ")


var appUser user

appUser = user{
	firstName: userfirstName, 
	lastName: userlastName,
	birthdate: userbirthdate, 
	createdAt: time.Now(),
}

	fmt.Println(userlastName,userfirstName,userbirthdate)
}

func getUserData(promptText string) string{
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}