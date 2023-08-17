package main

import (
	"basic/chatroom/client/processor"
	"fmt"
)

const LogIn = 1
const SignUp = 2
const LogOut = 3

type menuView struct {
	key  int
	loop bool
}

var username string
var password string

func (mv *menuView) mainMenu() {
	for mv.loop {
		fmt.Println("===============MENU==================")
		fmt.Println("1.Log in")
		fmt.Println("2.Sign up")
		fmt.Println("3.Log out")
		fmt.Println("Please Select 1-3:")

		_, err := fmt.Scanln(&mv.key)
		if err != nil {
			fmt.Println("Error Scan")
			return
		}

		switch mv.key {
		case LogIn:
			fmt.Println("Hello Log in")
			fmt.Println("username:")
			fmt.Scanln(&username)
			fmt.Println("password:")
			fmt.Scanln(&password)

			up := &processor.UserProcessor{}
			up.Login(username, password)
		case SignUp:
			fmt.Println("Hello Log up")
			fmt.Println("username:")
			fmt.Scanln(&username)
			fmt.Println("password:")
			fmt.Scanln(&password)

			up := &processor.UserProcessor{}
			up.Register(username, password)
		case LogOut:
			fmt.Println("Hello Log out")
		default:
			fmt.Println("False Input")
		}

	}
}

func main() {
	menuView := &menuView{key: 0, loop: true}
	menuView.mainMenu()
}
