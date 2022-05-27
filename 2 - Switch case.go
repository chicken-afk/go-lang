package main

import "fmt"

func main() {
	name := "Yuda Andika"
	switch name {
	case "Andika":
		fmt.Println("Hallo Andika")
		fmt.Println("Salam Kenal " + name)
	case "Yuda":
		fmt.Println("Hallo Yuda")
		fmt.Println("Hajimemaste " + name)
	default:
		fmt.Println("Siapa ya?")
	}

	// SWITCH DENGAN SHORT STATEMENT
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu Pendek")
	}

}
