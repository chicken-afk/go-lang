package main

import "fmt"

func main() {
	var name1 = "Andika"
	var name2 = "Andik"

	if name1 != name2 {
		fmt.Println("Tidak Sama")
	}

	if name1 == "Andik" {
		fmt.Println("masuk if")
	} else if name1 == "Andika" && name2 == "Andika" {
		fmt.Println("masuk else if")
	} else {
		fmt.Println("masuk ke else")
	}

	// var result bool = name1 != name2
	// fmt.Println(result)

	var (
		number1 = 100
		number2 = 150
	)
	fmt.Println(number1 > number2)

	// IF DENGAN SHORT STATEMENT
	if length := len(name1); length > 0 {
		fmt.Println("Lenght lebih dari 0")
	} else {
		fmt.Println("Length kurang dari 0")
	}

	var usia int8 = 35

	switch {
	case usia < 10:
		fmt.Print("Masih muda")
	case usia < 20:
		fmt.Print("Remaja")
	case usia < 30:
		fmt.Print("Dewasa")
	default:
		fmt.Print("Tua")
	}

}
