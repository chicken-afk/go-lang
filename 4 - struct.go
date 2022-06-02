package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var CustomerPremium Customer
	CustomerPremium.Name = "Andika Yuda"
	CustomerPremium.Address = "Shibuya, Japan"
	CustomerPremium.Age = 25
	fmt.Println(CustomerPremium)
	fmt.Println(CustomerPremium.Name)
}
