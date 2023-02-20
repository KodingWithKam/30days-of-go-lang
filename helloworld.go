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
	const BIRTH_YEAR = 1995
	const CURRENT_YEAR = 2023

	// Simply prints out contents when run
	fmt.Println("Constants", BIRTH_YEAR, CURRENT_YEAR)
}
