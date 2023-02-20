/*
A Go file consists of:

Package declaration
Imported packages
Functions
Statements and expressions
*/

package main

import (
	"fmt"
)

func main() {
	/*
		const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

		RULES FOR CONSTANTS:
		* Constant names follow the same naming rules as variables
		* Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
		* Constants can be declared both inside and outside of a function
	*/

	// untyped constant
	const BirthYear = 1995

	// typed constant
	const CurrentYear int = 2023

	// Multiple constants declaration
	const (
		A int = 5
		B     = 21
		C     = "Hi!"
	)

	// Simply prints out contents when run
	fmt.Println("Constants", BirthYear, CurrentYear, A, B, C)
}
