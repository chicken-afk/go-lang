package main

import "fmt"

func main() {
	//Tipe data array
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
	}
	var slice1 = months[0:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	months[1] = "ubah"
	fmt.Println(months)
	fmt.Println(slice1)
	slice1[1] = "ubahslice"
	slice1 = append(slice1, "dataappend")
	fmt.Println(months)
	fmt.Println(slice1)

}
