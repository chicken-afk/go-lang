package main

import "fmt"

func main() {

	var name = "Andika"
	fmt.Println("Hello World")
	fmt.Println("Nama saya ", name)

	kalimat := "nama saya adalah " + name
	fmt.Println(kalimat)

	gender := "Laki-laki"
	fmt.Println(gender)

	var (
		firstName = "Andika"
		// lastName  = "Yuda"
	)

	fmt.Println("Name Lengkap : ", firstName)
}
