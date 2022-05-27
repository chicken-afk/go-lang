package main

import "fmt"

func main() {
	type NoKTP string
	type Email string
	type Married bool

	var name NoKTP = "993893"
	var mail Email = "andikayuda628@gmail.com"
	var marriedStatus Married = true

	fmt.Println(name)
	fmt.Println(mail)
	fmt.Println(marriedStatus)
}
