package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Andika"
	names[1] = "Prasetya"
	names[2] = "Putra"

	fmt.Println(names)
	fmt.Println(names[1])

	//Array Langsung

	var values = []int{
		89,
		90,
		100,
	}

	fmt.Println(values)
	fmt.Println(values[1])

	//array function
	fmt.Println(len(values))

	values[2] = 72
	fmt.Println(values)

	var lagi [5]string
	fmt.Println(len(lagi))
}
