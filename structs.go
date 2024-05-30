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

func (u * user) outputUserDetails (){
fmt.Println(u.firstName,u.lastName,u.birthdate)
}

func (u *user) clearUserName(){
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName,birthdate string) *user{
	return &user{
	firstName: firstName, 
	lastName: lastName,
	birthdate: birthdate, 
	createdAt: time.Now(),
}
}

func main (){
	userfirstName :=getUserData("Please enter your first name: ")
	userlastName :=getUserData("Please enter your last name: ")
	userbirthdate :=getUserData("Please enter your bithdate (MM/DD/YYYY): ")


  var appUser *user

  appUser = newUser(userlastName,userfirstName,userbirthdate)

 appUser.outputUserDetails()
 appUser.clearUserName()
}

func getUserData(promptText string) string{
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}