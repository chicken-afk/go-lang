package main

import "fmt"

func main() {
	var nilai32 int32 = 2322
	var nilai64 = int64(nilai32)
	var nilai16 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Andika Yuda"
	var e = name[2]
	var eString string = string(e)
	fmt.Println(eString)

	names := "Rose"
	fmt.Println(string(names[0]))
}
