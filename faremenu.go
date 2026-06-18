package main

import "fmt"

func main() {
	var origin, destination, cabinClass string

	fmt.Println("*** Silver Karma Airfare Calculator")
	fmt.Print("Enter origin code: ")
	fmt.Scanln(&origin)
	fmt.Println("You've entered " + origin)

	fmt.Print("Enter destination code: ")
	fmt.Scanln(&destination)
	fmt.Println("You've entered " + destination)

	fmt.Print("Enter cabin class code: ")
	fmt.Scanln(&cabinClass)
	fmt.Println("You've entered " + cabinClass)
}
