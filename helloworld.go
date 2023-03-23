/*
A Go file consists of:

Package declaration
Imported packages
Functions
Statements and expressions
*/

package main

import "fmt"

func main() {
	/*
		GO OPERATORS 😎
		operators are used to perform operations on variables and values


		ARITHMETIC OPERATORS 🧮
		+	Addition	Adds together two values	x + y
		-	Subtraction	Subtracts one value from another	x - y
		*	Multiplication	Multiplies two values	x * y
		/	Division	Divides one value by another	x / y
		%	Modulus	Returns the division remainder	x % y
		++	Increment	Increases the value of a variable by 1	x++
		--	Decrement	Decreases the value of a variable by 1	x--

		ASSIGNMENT OPERATORS 🤨
		=	x = 5	x = 5
		+=	x += 3	x = x + 3
		-=	x -= 3	x = x - 3
		*=	x *= 3	x = x * 3
		/=	x /= 3	x = x / 3
		%=	x %= 3	x = x % 3
		&=	x &= 3	x = x & 3
		|=	x |= 3	x = x | 3
		^=	x ^= 3	x = x ^ 3
		>>=	x >>= 3	x = x >> 3
		<<=	x <<= 3	x = x << 3

		COMPARISON OPERATORS 😱
		==	Equal to	x == y
		!=	Not equal	x != y
		>	Greater than	x > y
		<	Less than	x < y
		>=	Greater than or equal to	x >= y
		<=	Less than or equal to	x <= y
	*/

	// Arithmetic Examples
	var x int = 4
	var y int = 3
	var z int = 2

	fmt.Print(x + y) // 4 + 3 = 7 ✅
	fmt.Print(x - y) // 4 - 3 = 1 ✅
	fmt.Print(x * y) // 4 * 3 = 12 ✅
	fmt.Print(x / z) // 4 / 2 = 2 ✅
	fmt.Print(x % z) // 4 % 2 = 0 ✅

	// ASSIGNMENT OPERATORS EXAMPLES
	var a int = 1
	var b int = 5
	a += 3
	fmt.Print(a) // 4 ✅

	a -= 3
	fmt.Print(a) // 1 ✅

	a *= 3
	fmt.Print(a) // 3 ✅

	a /= 3
	fmt.Print(a) // 1 ✅

	a %= 1
	fmt.Print(a) // 0 ✅

	b &= 3
	fmt.Printf("b is %b \n", b)   // 101 ✅
	fmt.Printf("3 is %03b \n", 3) // 011 ✅

	b &= 3
	fmt.Printf("b now is %03b \n", b) // 001 ✅

	b |= 3
	fmt.Printf("b now is %03b \n", b) // 111 ✅

	b ^= 3
	fmt.Printf("b now is %03b \n", b) // 110 ✅

	b >>= 3
	fmt.Printf("b now is %03b \n", b) // 000 ✅

	b <<= 3

	fmt.Printf("b now is %03b \n", b) // 101000 ✅

	// COMPARISON OPERATORS EXAMPLES
	var c int = 6

	fmt.Print(c != 7) // 6 not equal to 7 true ✅
	fmt.Print(c > 7)  // 6 > 7 false ✅
	fmt.Print(c < 7)  // 6 > 7 true ✅
	fmt.Print(c <= 7) // 6 less than or equal to 7 true ✅
	fmt.Print(c >= 7) // 6 greater than or equal to 7 false ✅
}
