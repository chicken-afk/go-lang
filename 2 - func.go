package main

import "fmt"

func kali(a int, b int) int {
	hasil := a * b
	hasil = int(hasil)
	return hasil
}

// Returning multiple values

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

// Function return named value

func getFullName2() (firstName string, middleName string, usia int) {
	firstName = "Andika"
	middleName = "Yuda"
	usia = 28

	return
}

func main() {
	hasil := kali(2, 3)
	fmt.Println(hasil)

	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// Ignore return value

	firstname, _ := getFullName()
	fmt.Println(firstname)
	// fmt.Println(lastName)

	a, b, usia := getFullName2()

	fmt.Println("Nama : ", a, b, ".Usia : ", usia)
}
