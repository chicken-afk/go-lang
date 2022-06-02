package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

type Data []struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	var CustomerPremium Customer
	CustomerPremium.Name = "Andika Yuda"
	CustomerPremium.Address = "Shibuya, Japan"
	CustomerPremium.Age = 25
	fmt.Println(CustomerPremium)
	fmt.Println(CustomerPremium.Name)

	// Struct Literals

	CustomerReguler := Customer{
		Name:    "Prasety Putra",
		Address: "New York City",
		Age:     33,
	}
	fmt.Println(CustomerReguler)

	CustomerFirstClass := Customer{"Gojo Satoru", "Shibuya", 28}
	fmt.Println(CustomerFirstClass)

	// One Dimentional struct
	d := Data{
		{"message1", "one"}, {"message2", "two"},
		{"message3", "three"}, {"message4", "four"},
	}
	fmt.Println(d)
	fmt.Println(d[0])
	fmt.Println(d[1].Message)
}
