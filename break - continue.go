package main

import "fmt"

func main() {
	// Break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("")
	// Continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
