package main

import (
	"fmt"
	"os"
	"strings"
)

/*
*
Reordering Names Based on Country Code

Goal: Create a program that reorders names
based on the specified country's format.

Inputs: First name, last name, middle name
(optional), and country code.

Outputs: Reordered name based on the
country's format
*
*/
func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("missing arguments")
	}

	countryCode := strings.ToUpper(args[len(args)-1])
	firstName := args[0]
	lastName := args[1]
	middleNames := args[2 : len(args)-1]

	firstLastCountries := []string{"US"} // The list can be extended
	lastFirstCountries := []string{"VN"} // The list can be extended

	switch {
	case isAvailable(firstLastCountries, countryCode):
		// Combine the name by this format: firstName, middleNames, lastName
		printWithSpaces(append(append([]string{firstName}, middleNames...), lastName))
	case isAvailable(lastFirstCountries, countryCode):
		// Combine the name by this format: lastName, middleNames, firstName
		printWithSpaces(append(append([]string{lastName}, middleNames...), firstName))
	default:
		fmt.Println("Unsupported country code")
	}
}

func printWithSpaces(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s ", arr[i])
	}
}

func isAvailable(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
	}
	return false
}
