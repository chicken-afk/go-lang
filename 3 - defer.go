package main

import "fmt"

func logging(name string) {
	fmt.Println("Finish start logging system", name)
}

func runApplication() {
	/** Even if in runApplication run error the logging function still run*/
	defer logging("Andika")
	fmt.Println("Run System")
}

func main() {
	runApplication()
}
