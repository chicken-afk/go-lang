package main

import "fmt"

func kali(a int, b int) int {
	hasil := a * b
	hasil = int(hasil)
	return hasil
}

func main() {
	hasil := kali(2, 3)
	fmt.Println(hasil)
}
