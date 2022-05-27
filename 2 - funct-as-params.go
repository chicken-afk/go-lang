package main

import "fmt"

func sayHelloWithFillter(name string, fillter func(string) string) {
	nameFilltered := fillter(name)
	fmt.Println("Hello ", nameFilltered)
}

func spamFillter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFillter("Anjing", spamFillter)
}
