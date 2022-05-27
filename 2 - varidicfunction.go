package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	hasilTambah := sumAll(10, 5, 7, 90, 100)
	fmt.Println("Hasil : ", hasilTambah)

	// hasil := sumAll(5, 5, 78, 90, 100)
	fmt.Println("Hasil :", sumAll(5, 5, 5, 90, 100))

	// Initiate slice
	numberSlicec := []int{10, 10, 10, 10, 5}
	total := sumAll(numberSlicec...)
	fmt.Println("Hasil Total : ", total)
}
