package main

import "fmt"

func main() {
	i := 1
	a := 98
	for i <= 100 {
		fmt.Println("Perulangan ke ", i+a)
		i++
	}

	// For Dengan Statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// For Range
	slice := []string{
		"Andika",
		"Yuda",
		"Prasetya",
		"Putra",
	}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Eko", "Kurniawan", "Dadang"}
	for key, value := range names {
		fmt.Println("Index ke ", key, "= ", value)
	}

	// Jika tidak ingin memakai key ganti key dengan underscore _
	for _, value := range names {
		fmt.Println(value)
	}

	person := map[string]string{
		"name":  "Andika",
		"title": "System Analyst",
	}

	for key, value := range person {
		fmt.Println(key, " : ", value)
	}
}
