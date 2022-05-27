package main

import "fmt"

func main() {
	var a = 10
	var b = 2
	a++

	tambah := a + b
	kali := a * b

	fmt.Println(tambah)
	fmt.Println(kali)

	a += 10 // a = a+10
	b -= 2
	fmt.Println(a)
	fmt.Println(b)
}
