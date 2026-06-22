package main

import (
	"errors"
	"fmt"
	"strings"
)

func getCityFromCode(codeToLookup string) (City, error) {
	for _, item := range cities {
		if item.code == strings.ToUpper(codeToLookup) {
			return item, nil
		}
	}

	message := fmt.Sprintf("%s is not a valid city code", codeToLookup)

	return City{"", "", -999, -999}, errors.New(message)
}

func getCabinClassFromCode(codeToLookup string) (CabinClass, error) {
	for _, item := range cabinClasses {
		if item.code == strings.ToUpper(codeToLookup) {
			return item, nil
		}
	}

	message := fmt.Sprintf("%s is not a valid cabin class code", codeToLookup)

	return CabinClass{"", "", -999}, errors.New(message)
}
