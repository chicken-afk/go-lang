package main

import "fmt"

func main() {
	users := map[string]string{
		"name": "Andika",
		"usia": "25",
	}

	users["alamat"] = "Kebumen"

	fmt.Println(users)
	fmt.Println(users["name"])
	fmt.Println(users["usia"])

	bank := make(map[string]string)
	bank["account_number"] = "12312744"
	bank["account_type"] = "BCA"
	fmt.Println(bank)
	fmt.Println(bank["account_number"])

	delete(bank, "account_type")
	fmt.Println(bank["account_number"])
	fmt.Println(len(bank))
}
