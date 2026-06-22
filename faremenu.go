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

		for i := range cabinClasses {
			fmt.Println(cabinClasses[i].code + " = " + cabinClasses[i].className)
		}

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

	distance := calculateDistance(float64(destinationCity.longitude)/10000,
		float64(destinationCity.latitude)/10000,
		float64(originCity.longitude)/10000,
		float64(originCity.latitude)/10000)

	fmt.Printf("\nDistance: %.1f km\n", distance)
	rateFloat32 := float32(enteredCabinClass.rate) / 100

	fmt.Printf("$ per kilometer: %.2f\n", rateFloat32)
	fmt.Printf("Total fare: $%.2f\n", rateFloat32*float32(distance))
}
