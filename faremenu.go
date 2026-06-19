package main

import "fmt"

func main() {
	var origin, destination, cabinClass string

	validOriginEntered := false
	var originCity City
	var originErr error

	fmt.Println("*** Silver Karma Airfare Calculator")

	for !validOriginEntered {
		fmt.Print("Enter origin code: ")
		fmt.Scanln(&origin)

		originCity, originErr = getCityFromCode(origin)

		if originErr == nil {
			fmt.Println("You've entered " + originCity.cityName)
			validOriginEntered = true
		} else {
			fmt.Println(originErr)
		}
	}

	fmt.Print("Enter destination code: ")
	fmt.Scanln(&destination)
	fmt.Println("You've entered " + destination)

	fmt.Print("Enter cabin class code: ")
	fmt.Scanln(&cabinClass)
	fmt.Println("You've entered " + cabinClass)

}
