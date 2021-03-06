package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Login berhasil", name)
	}
}

func main() {
	// anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("admin", blacklist)
	registerUser("anjing", blacklist)

}
