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

func main() {
	hasil := kali(2, 3)
	fmt.Println(hasil)

	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
