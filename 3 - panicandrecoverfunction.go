package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error message : ", message)
	} else {
		fmt.Println("Apps run succesfully")
	}

	fmt.Println("End of Application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Internal server error")
	}
	fmt.Println("Application running...")
}

func main() {
	runApp(false)
}
