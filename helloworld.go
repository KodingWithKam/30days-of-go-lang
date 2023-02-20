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
			GO has 3 functions to output text

			Print() --> prints its arguments with their default format

			Println() --> similar to Print() with the difference that a whitespace is added between the arguments,
			and a newline is added at the end

			Printf() --> first formats its argument based on the given formatting verb and then prints them.
		    	*FORMATTING VERBS BELOW*
				%v is used to print the value of the arguments
				%T is used to print the type of the arguments
				%#v	Prints the value in Go-syntax format
				%%	Prints the % sign

				*INTEGER FORMATTING VERBS*
				%b	Base 2
				%d	Base 10
				%+d	Base 10 and always show sign
				%o	Base 8
				%O	Base 8, with leading 0o
				%x	Base 16, lowercase
				%X	Base 16, uppercase
				%#x	Base 16, with leading 0x
				%4d	Pad with spaces (width 4, right justified)
				%-4d	Pad with spaces (width 4, left justified)
				%04d	Pad with zeroes (width 4)

				*STRING FORMATTING VERBS*
				%s	Prints the value as plain string
				%q	Prints the value as a double-quoted string
				%8s	Prints the value as plain string (width 8, right justified)
				%-8s	Prints the value as plain string (width 8, left justified)
				%x	Prints the value as hex dump of byte values
				% x	Prints the value as hex dump with spaces

				*BOOLEAN FORMATTING VERBS*
				%t	Value of the boolean operator in true or false format (same as using %v)

				*FLOATING POINT FORMATTING VERBS*
				%e	Scientific notation with 'e' as exponent
				%f	Decimal point, no exponent
				%.2f	Default width, precision 2
				%6.2f	Width 6, precision 2
				%g	Exponent as needed, only necessary digits

	*/
	var i, j, k string = "Hello", "World", "Kam"
	var x, y, z int = 1, 3, 15
	var isBool bool = false
	const PI = 3.14

	fmt.Print(i)
	fmt.Print(j)

	// Print with space between both strings
	fmt.Print(i, " ", j) // output should be Hello World

	// Print() inserts a space between the arguments if neither values are strings
	fmt.Print(x, y) // output should be 10 20

	fmt.Println(i, j) // output should be Hello World

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("x has value: %v and type: %T", x, x) //output x has value: 1 and type: int
	fmt.Printf("%v%%\n", x)                          // output 1%

	// INTEGER FORMATTING EXAMPLES
	fmt.Printf("%b\n", z)   // output 1111
	fmt.Printf("%d\n", z)   // output 15
	fmt.Printf("%+d\n", z)  // output +15
	fmt.Printf("%o\n", z)   // output 17
	fmt.Printf("%O\n", z)   // output 0o17
	fmt.Printf("%x\n", z)   // output f
	fmt.Printf("%X\n", z)   // output F
	fmt.Printf("%#x\n", z)  // output 0xf
	fmt.Printf("%4d\n", z)  // output   15
	fmt.Printf("%-4d\n", z) // output 15
	fmt.Printf("%04d\n", z) // output 0015

	// STRING FORMATTING EXAMPLES
	fmt.Printf("%s\n", k)   // output Kam
	fmt.Printf("%q\n", k)   // output "Kam"
	fmt.Printf("%8s\n", k)  // output      Kam
	fmt.Printf("%-8s\n", k) // output Kam
	fmt.Printf("%x\n", k)   // output 4b616d
	fmt.Printf("% x\n", k)  // output 4b 61 6d

	// BOOLEAN FORMATTING EXAMPLES
	fmt.Printf("%t\n", isBool)

	// FLOATING POINT FORMATTING EXAMPLES
	fmt.Printf("%e\n", PI)    //output 3.141000e+00
	fmt.Printf("%f\n", PI)    //output 3.141000
	fmt.Printf("%.2f\n", PI)  //output 3.14
	fmt.Printf("%6.2f\n", PI) //output   3.14
	fmt.Printf("%g\n", PI)    //output 3.141
}
