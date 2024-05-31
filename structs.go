package main

import "fmt"

import "example.com/structs/user"

func main (){
	userfirstName :=getUserData("Please enter your first name: ")
	userlastName :=getUserData("Please enter your last name: ")
	userbirthdate :=getUserData("Please enter your bithdate (MM/DD/YYYY): ")


  var appUser *user.User

  appUser,err := user.NewUser(userlastName,userfirstName,userbirthdate)

 if err != nil {
	fmt.Println(err)
	return
 }

 admin := user.NewAdmin("test@example.com","test123")
 admin.ClearUserName()
 admin.OutputUserDetails()
 

 appUser.OutputUserDetails()
 appUser.ClearUserName()
 appUser.OutputUserDetails()
}

func getUserData(promptText string) string{
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}