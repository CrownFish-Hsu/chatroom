package main

import (
	"basic/chatroom/client/service"
	"fmt"
)

const LogIn = 1
const SignUp = 2
const LogOut = 3

type menuView struct {
	key         int
	loop        bool
	userService *service.UserService
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
			mv.loop = false
		case SignUp:
			fmt.Println("Hello Log up")
			mv.loop = false
		case LogOut:
			fmt.Println("Hello Log out")
			mv.loop = false
		default:
			fmt.Println("False Input")
		}

		if !mv.loop {
			break
		}
	}

	if LogIn == mv.key {
		fmt.Println("username:")
		fmt.Scanln(&username)
		fmt.Println("password:")
		fmt.Scanln(&password)

		mv.userService.Login(username, password)
	}
}

func main() {
	menuView := &menuView{key: 0, loop: true}
	menuView.mainMenu()
}
