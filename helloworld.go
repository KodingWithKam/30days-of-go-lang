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

/*
In Go, there are different types of variables:

1. int- stores integers (whole numbers), such as 123 or -123
2. float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
3. string - stores text, such as "Hello World". String values are surrounded by double quotes
4. bool- stores values with two states: true or false
*/

func main() {
	/*
		DECLARING VARIABLES IN GO..

		IN GO THERE ARE 2 WAYS TO DECLARE VARIABLES

		1. Use the var keyword, followed by variable name and type (ie, int, string, etc)
		ex. var variableName type = value

		2. Using the := sign
		ex. variableName := value

		var VS. :=
		- var can be used inside and outside functions while := can only be used inside functions
		- var can be declared separately (ie var firstName string)
		- := cannot be declared separately must be done in the same line

		NAMING VARIABLES IN GO!
		- A variable can have a short name (like x and y) or a more descriptive name (age, lastName, etc.)

		Go variable naming rules:

		- A variable name must start with a letter or an underscore character (_) ** cannot start with a number!
		- A variable name can only contain alphanumeric characters and underscores (a-z, A-Z, 0-9, and _ )
		- Variable names are case-sensitive (age, Age and AGE are three different variables)
		- There is no limit on the length of the variable name
		- A variable name cannot contain spaces
		- The variable name cannot be any Go keywords

		NAMING CONVENTIONS IN GO

		- Camel Case
		ex. myVariableName = "John"

		- Pascal Case
		ex. MyVariableName = "John"

		- Snake Case
		ex. my_variable_name = "John"

	*/

	var isBoolean bool = true
	var isFalse bool = false

	// Naming Techniques 4 Variables
	var is_snake_case = true // <-- ✅EVERY WORD SEPARATED WITH UNDERSCORE

	var IsPascalCase = true // <-- ✅EVERY WORD BEGINS WITH CAPITAL LETTER

	var isCamelCase = true // <-- ✅EVERY WORD EXCEPT 1st BEGINS WITH CAPITAL LETTER

	// Naming Conventions
	var _startsWithUnderScore = true      // <-- ✅THIS IS OK!
	var startsWithLetter string           // <-- ✅THIS IS OK!
	var superDuperLongVariableName string // <-- ✅THIS IS OK!
	var b = "B"                           // <-- ✅THIS IS OK! (CASE SENSITIVE)
	var B = "B"                           // <-- ✅THIS IS OK! (CASE SENSITIVE)
	// var 1startsWithNumber = 1 // <-- ❌THIS IS NOT OK!
	// var var = 1 // <-- ❌THIS IS NOT OK!
	// var name with spaces = "long" // <-- ❌THIS IS NOT OK!

	// var vs :=
	var newName string                       // <-- we can add a value later!
	newName = "Declared After variable init" // <-- adding value here!

	y := 10 // <-- We NEED TO SET THE VALUE IN THE SAME LINE!

	// Example 1 with var keyword
	var a = 25
	var firstName string = "Kam"
	var lastName string = "Haze"
	var middleName string

	// Example 2 with := symbol (notice how var keyword is omitted)
	x := 2
	test := "HELLO"

	// You can also declare variables without setting the value if it's not know
	var c bool

	// You can set the above value here if needed
	c = false

	// Multiple Variable Declaration on same line
	var first, second, third, fourth string = "first", "second", "third", "fourth"

	// Print out above examples
	fmt.Println(firstName)
	fmt.Println(x, y, a)
	fmt.Println(c)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(lastName)
	fmt.Println(middleName, test, newName, startsWithLetter, _startsWithUnderScore,
		superDuperLongVariableName, b, B, isCamelCase, IsPascalCase, is_snake_case,
		isBoolean, isFalse,
	)

	// Simply prints out Hello World when program is run
	fmt.Println("Hello World!")
}
