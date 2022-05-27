package main

import "fmt"

type Fillter func(string) string

func sayHelloWithFillter(name string, fillter func(string) string) {
	nameFilltered := fillter(name)
	fmt.Println("Hello ", nameFilltered)
}

func sayHelloWithFillterType(name string, filter Fillter) {
	nameFilltered := filter(name)
	fmt.Println("Hello ", nameFilltered)
}

func spamFillter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

// Function type declaration

func main() {
	sayHelloWithFillter("Anjing", spamFillter)
	sayHelloWithFillter("Andika", spamFillter)
}
