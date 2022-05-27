package main

import "fmt"

func main() {
	getGoodBye := getGoodBye
	fmt.Println(getGoodBye("Andika"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
