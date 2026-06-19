package main

import "fmt"

func main() {
	var origin, destination, cabinClass string

	fmt.Println("*** Silver Karma Airfare Calculator")

	validOriginEntered := false
	var originCity City
	var originErr error

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

	validDestinationEntered := false
	var destinationCity City
	var destinationErr error

	for !validDestinationEntered {
		fmt.Print("Enter destination code: ")
		fmt.Scanln(&destination)

		destinationCity, destinationErr = getCityFromCode(destination)

		if destinationErr == nil {
			fmt.Println("You've entered " + destinationCity.cityName)
			validDestinationEntered = true
		} else {
			fmt.Println(destinationErr)
		}
	}

	validCabinClassEntered := false
	var enteredCabinClass CabinClass
	var enteredCabinClassErr error

	for !validCabinClassEntered {

		fmt.Print("Enter cabin class code: ")
		fmt.Scanln(&cabinClass)

		enteredCabinClass, enteredCabinClassErr = getCabinClassFromCode(cabinClass)

		if enteredCabinClassErr == nil {
			fmt.Println("You've entered " + enteredCabinClass.className + " class")
			validCabinClassEntered = true
		} else {
			fmt.Println(enteredCabinClassErr)
		}
	}

}
